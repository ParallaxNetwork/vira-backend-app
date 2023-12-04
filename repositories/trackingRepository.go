package repositories

import (
	"fmt"
	"vira-backend-app/enums"
	"vira-backend-app/models"
	"vira-backend-app/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TrackingRepository interface {
	FindById(id string) (*models.Tracking, error)
	FindAll() ([]models.Tracking, error)
	InsertOne(tracking models.Tracking) (*models.Tracking, error)
	DeleteById(id string) (error)
	UpdateById(id string, updatedTracking *models.Tracking) (*models.Tracking, error)
}

type trackingRepository struct{}

func NewTrackingRepository() TrackingRepository {
	return &trackingRepository{}
}

func (r *trackingRepository) FindById(id string) (*models.Tracking, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var tracking models.Tracking

	filter := bson.M{"_id": id}

	err = db.Collection("trackings").FindOne(ctx, filter).Decode(&tracking)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("No Tracking Found: %v", err)
		}

		return nil, fmt.Errorf("Database Error: %v", err)
	}

	return &tracking, nil
}

func (r *trackingRepository) FindAll() ([]models.Tracking, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var trackings []models.Tracking

	cursor, err := db.Collection("trackings").Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("Database Error: %v", err)
	}

	err = cursor.All(ctx, &trackings)
	if err != nil {
		return nil, fmt.Errorf("Database Error: %v", err)
	}

	return trackings, nil
}

func (r *trackingRepository) InsertOne(tracking models.Tracking) (*models.Tracking, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	tracking.TrackingId = utils.GenerateRandomID()
	tracking.Status 		= enums.Ongoing

	_, err = db.Collection("trackings").InsertOne(ctx, tracking)
	if err != nil {
		return nil, fmt.Errorf("Database Error: %v", err)
	}

	return &tracking, nil
}

func (r *trackingRepository) DeleteById(id string) (error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	filter := bson.M{"_id": id}
	_, err = db.Collection("trackings").DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func (r *trackingRepository) UpdateById(id string, updatedTracking *models.Tracking) (*models.Tracking, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedTracking}

	_, err = db.Collection("trackings").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return updatedTracking, nil
}
