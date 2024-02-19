package repositories

import (
	"context"
	"fmt"
	"math/big"

	"vira-backend-app/dto"
	"vira-backend-app/models"
	"vira-backend-app/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FundingRepository interface {
	FindById(id string) (*models.Funding, error)
	FindAll() ([]models.Funding, error)
	FindAllByProjectId(projectId string) ([]models.Funding, error)
	FindAllByUserId(userId string) ([]models.Funding, error)
	InsertOne(funding dto.FundingInsertOneDTO) (*models.Funding, error)
	DeleteById(id string) error
	UpdateById(id string, updatedFunding *models.Funding) (*models.Funding, error)
}

type fundingRepository struct{}

func NewFundingRepository() FundingRepository {
	return &fundingRepository{}
}

func (r *fundingRepository) FindById(id string) (*models.Funding, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var funding models.Funding

	filter := bson.M{"_id": id}

	err = db.Collection("fundings").FindOne(ctx, filter).Decode(&funding)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no funding found: %v", err)
		}

		return nil, fmt.Errorf("database error: %v", err)
	}

	return &funding, nil
}

func (r *fundingRepository) FindAll() ([]models.Funding, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var fundings []models.Funding

	cursor, err := db.Collection("fundings").Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf(err.Error())

	}

	err = cursor.All(ctx, &fundings)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return fundings, nil
}

func (r *fundingRepository) FindAllByProjectId(projectId string) ([]models.Funding, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var fundings []models.Funding

	filter := bson.M{"projectId": projectId}

	cursor, err := db.Collection("fundings").Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	err = cursor.All(ctx, &fundings)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return fundings, nil
}

func (r *fundingRepository) FindAllByUserId(userId string) ([]models.Funding, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var fundings []models.Funding

	filter := bson.M{"user_id": userId}

	cursor, err := db.Collection("fundings").Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	err = cursor.All(ctx, &fundings)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return fundings, nil
}

func (r *fundingRepository) InsertOne(fundingDTO dto.FundingInsertOneDTO) (*models.Funding, error) {
	// Get contract, wallet and databases ready
	contract, client := utils.ContractConnect()
	wallet := utils.WalletConnect(client)
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	// Find current funding with matching userId and projectId
	var currFunding models.Funding
	_ = db.Collection("fundings").FindOne(ctx, bson.M{"user_id": fundingDTO.UserId, "project_id": fundingDTO.ProjectId}).Decode(&currFunding)

	// Find target project
	var project models.Project
	err = db.Collection("projects").FindOne(ctx, bson.M{"_id": fundingDTO.ProjectId}).Decode(&project)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	// Search user wallet as target wallet for minting
	var user models.User
	err = db.Collection("users").FindOne(ctx, bson.M{"_id": fundingDTO.UserId}).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var txn *types.Transaction
	if currFunding.FundingId != "" {
		// Add new funding to current tokenId
		hexString := currFunding.TokenId[2:]
		currFundTokenId, _ := new(big.Int).SetString(hexString, 16)
		txn, err = contract.MintValue(wallet, currFundTokenId, big.NewInt(int64(fundingDTO.Amount*1e6)))
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		bind.WaitMined(context.Background(), client, txn)
	} else {
		// Mint new token to user wallet from available campaign
		txn, err = contract.Mint(wallet, common.HexToAddress(user.WalletAddress), fundingDTO.ProjectId, big.NewInt(int64(fundingDTO.Amount*1e6)))
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		bind.WaitMined(context.Background(), client, txn)
	}

	// Get token ID from recent txn
	txnReceipt, err := client.TransactionReceipt(context.Background(), txn.Hash())
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	// Update funding data
	if currFunding.FundingId == "" {
		var funding models.Funding
		funding.FundingId = utils.GenerateRandomID()
		funding.ProjectId = fundingDTO.ProjectId
		funding.UserId = fundingDTO.UserId
		funding.TransactionHistory = append(funding.TransactionHistory, models.FundingTransaction{
			Amount:   float64(fundingDTO.Amount),
			AdminFee: float64(fundingDTO.AdminFee),
			Total:    float64(fundingDTO.Total),
		})
		funding.TokenId = txnReceipt.Logs[4].Topics[2].String()
		funding.CreatedAt = utils.GetCurrentTime()

		_, err = db.Collection("fundings").InsertOne(ctx, funding)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}

		return &funding, nil
	} else {
		currFunding.TransactionHistory = append(currFunding.TransactionHistory, models.FundingTransaction{
			Amount:   float64(fundingDTO.Amount),
			AdminFee: float64(fundingDTO.AdminFee),
			Total:    float64(fundingDTO.Total),
		})

		_, err = db.Collection("fundings").UpdateOne(ctx, bson.M{"_id": currFunding.FundingId}, bson.M{"$set": currFunding})
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}

		return &currFunding, nil
	}
}

func (r *fundingRepository) DeleteById(id string) error {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	_, err = db.Collection("fundings").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func (r *fundingRepository) UpdateById(id string, updatedFunding *models.Funding) (*models.Funding, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedFunding}

	_, err = db.Collection("fundings").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return updatedFunding, nil
}
