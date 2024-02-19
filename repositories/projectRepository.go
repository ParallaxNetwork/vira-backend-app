package repositories

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"vira-backend-app/enums"
	"vira-backend-app/models"
	"vira-backend-app/utils"

	"vira-backend-app/contracts"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectRepository interface {
	FindById(projectId string) (*models.Project, error)
	FindAll() ([]models.Project, error)
	FindByAdminId(adminId string) (*models.Project, error)
	InsertOne(project models.Project) (*models.Project, error)
	DeleteById(projectId string) error
	UpdateById(projectId string, updatedProject *models.Project) (*models.Project, error)
	StartProject(projectId string) error
	RevokeProject(projectId string) error
}

type projectRepository struct{}

func NewProjectRepository() ProjectRepository {
	return &projectRepository{}
}

func (r *projectRepository) FindById(projectId string) (*models.Project, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var project models.Project

	filter := bson.M{"_id": projectId}

	err = db.Collection("projects").FindOne(ctx, filter).Decode(&project)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("NO PROJECT FOUND: %v", err)
		}

		return nil, fmt.Errorf("DATABASE ERROR: %v", err)
	}

	cursor, err := db.Collection("fundings").Find(ctx, bson.M{"project_id": project.ProjectId})
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var fundings []models.Funding
	var fundingCountTotal float64 = 0
	err = cursor.All(ctx, &fundings)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	for _, funding := range fundings {
		for _, transaction := range funding.TransactionHistory {
			fundingCountTotal += transaction.Amount
		}
	}

	project.CollectedValue = fundingCountTotal

	return &project, nil
}

func (r *projectRepository) FindAll() ([]models.Project, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var projects []models.Project

	cursor, err := db.Collection("projects").Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	err = cursor.All(ctx, &projects)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	for i := range projects {
		cursor, err := db.Collection("fundings").Find(ctx, bson.M{"project_id": projects[i].ProjectId})
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}

		var fundings []models.Funding
		var fundingCountTotal float64 = 0
		err = cursor.All(ctx, &fundings)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}

		for _, funding := range fundings {
			for _, transaction := range funding.TransactionHistory {
				fundingCountTotal += transaction.Amount
			}
		}

		projects[i].CollectedValue = fundingCountTotal
	}

	return projects, nil
}

func (r *projectRepository) FindByAdminId(adminId string) (*models.Project, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var project models.Project

	filter := bson.M{"admin_id": adminId}

	err = db.Collection("projects").FindOne(ctx, filter).Decode(&project)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("NO PROJECT FOUND: %v", err)
		}

		return nil, fmt.Errorf("DATABASE ERROR: %v", err)
	}

	return &project, nil
}

func (r *projectRepository) InsertOne(project models.Project) (*models.Project, error) {
	// Get contract, wallet and databases ready
	contract, client := utils.ContractConnect()
	wallet := utils.WalletConnect(client)
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	// Fill project data
	project.ProjectId = utils.GenerateRandomID()
	project.CreatedAt = utils.GetCurrentTime()
	project.Status = enums.Crowdfunding

	// Insert project
	_, err = db.Collection("projects").InsertOne(ctx, project)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	// Create campaign
	campaignData := contracts.IViraCampaignsCampaign{
		Id:          project.ProjectId,
		Start:       big.NewInt(0),
		End:         big.NewInt(0),
		RelInv:      big.NewInt(0),
		Slice:       big.NewInt(0),
		PctPerSlice: big.NewInt(0),
		Locked:      false,
		Revoked:     false,
	}

	// Update campaign
	_, err = contract.UpdateCampaign(wallet, campaignData)
	if err != nil {
		log.Fatal(err)
	}

	return &project, nil
}

func (r *projectRepository) DeleteById(projectId string) error {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	filter := bson.M{"_id": projectId}
	_, err = db.Collection("projects").DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func (r *projectRepository) UpdateById(projectId string, updatedProject *models.Project) (*models.Project, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	filter := bson.M{"_id": projectId}
	update := bson.M{"$set": updatedProject}

	_, err = db.Collection("projects").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return updatedProject, nil
}

func (r *projectRepository) StartProject(projectId string) error {
	// Get contract, wallet and databases ready
	contract, client := utils.ContractConnect()
	wallet := utils.WalletConnect(client)
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	// Find target project
	var project models.Project
	filter := bson.M{"_id": projectId}
	err = db.Collection("projects").FindOne(ctx, filter).Decode(&project)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	// Set time variables for project
	currentTime := time.Now()
	campaignStart := currentTime.Unix()
	campaignEnd := currentTime.AddDate(0, int(project.ContractDuration), 0).Unix()

	// Update campaign data
	campaignData := contracts.IViraCampaignsCampaign{
		Id:          project.ProjectId,
		Start:       big.NewInt(campaignStart),
		End:         big.NewInt(campaignEnd),
		RelInv:      big.NewInt(0),
		Slice:       big.NewInt(project.ContractDuration / 6),
		PctPerSlice: big.NewInt(int64((project.ReturnPerYear * 100) / 2)),
		Locked:      true,
		Revoked:     false,
	}

	project.Status = enums.Ongoing

	_, err = contract.UpdateCampaign(wallet, campaignData)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	update := bson.M{"$set": project}
	_, err = db.Collection("projects").UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf(err.Error())
	}

	return nil
}

func (r *projectRepository) RevokeProject(projectId string) error {
	// Get contract, wallet and databases ready
	contract, client := utils.ContractConnect()
	wallet := utils.WalletConnect(client)
	_ = wallet
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	// Find target project
	var project models.Project
	filter := bson.M{"_id": projectId}
	err = db.Collection("projects").FindOne(ctx, filter).Decode(&project)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	// Find campaign data
	campaignData, err := contract.CampaignData(nil, project.ProjectId)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	// Update campaign data
	campaignData.Revoked = true
	project.Status = enums.Failed

	fmt.Println(campaignData)
	_, err = contract.UpdateCampaign(wallet, campaignData)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	update := bson.M{"$set": project}
	_, err = db.Collection("projects").UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}
