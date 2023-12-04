package repositories

import (
	"fmt"

	"vira-backend-app/models"
	"vira-backend-app/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FundingRepository interface {
	FindById(id string) (*models.Funding, error)
	FindAll() ([]models.Funding, error)
	FindAllByProjectId(projectId string) ([]models.Funding, error)
	FindAllByUserId(userId string) ([]models.Funding, error)
	InsertOne(funding models.Funding) (*models.Funding, error)
	DeleteById(id string) (error)
	UpdateById(id string, updatedFunding *models.Funding) (*models.Funding, error)
}

type fundingRepository struct {}

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
			return nil, fmt.Errorf("No Funding Found: %v", err)
		}

		return nil, fmt.Errorf("Database Error: %v", err)
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

func (r *fundingRepository) InsertOne(funding models.Funding) (*models.Funding, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	funding.FundingId = utils.GenerateRandomID()
	funding.CreatedAt = utils.GetCurrentTime()

	_, err = db.Collection("fundings").InsertOne(ctx, funding)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &funding, nil
}

func (r *fundingRepository) DeleteById(id string) (error) {
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