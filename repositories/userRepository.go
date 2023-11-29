package repositories

import (
	"fmt"

	"vira-backend-app/models"
	"vira-backend-app/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindById(userID string) (*models.User, error)
	InsertOne(user models.User) (*models.User, error)
	DeleteById(userId string) (error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindById(userId string) (*models.User, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var user models.User
	filter := bson.M{"_id": userId}

	err = db.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("No User Found: %v", err)
		}

		return nil, fmt.Errorf("Database Error: %v", err)
	}

	return &user, nil
}

func (r *userRepository) InsertOne(user models.User) (*models.User, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	user.UserId = utils.GenerateRandomID()
	user.CreatedAt = utils.GetCurrentTime()

	_, err = db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &user, nil
}

func (r *userRepository) DeleteById(userId string) (error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	filter := bson.M{"_id": userId}
	_, err = db.Collection("users").DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}