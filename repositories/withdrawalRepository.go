package repositories

import (
	"fmt"

	"vira-backend-app/models"
	"vira-backend-app/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type WithdrawalRepository interface {
	FindAll() ([]models.Withdrawal, error)
	FindById(id string) (*models.Withdrawal, error)
	FindByUserId(userId string) ([]models.Withdrawal, error)
	InsertOne(withdrawal models.Withdrawal) (*models.Withdrawal, error)
	UpdateById(id string, updatedWithdrawal *models.Withdrawal) (*models.Withdrawal, error)
	DeleteById(id string) error
}

type withdrawalRepository struct {}

func NewWithdrawalRepository() WithdrawalRepository {
	return &withdrawalRepository{}
}

func (r *withdrawalRepository) FindAll() ([]models.Withdrawal, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	var withdrawals []models.Withdrawal

	cursor, err := db.Collection("withdrawals").Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find withdrawals: %v", err)
	}

	if err = cursor.All(ctx, &withdrawals); err != nil {
		return nil, fmt.Errorf("failed to read withdrawals cursor: %v", err)
	}

	return withdrawals, nil
}

func (r *withdrawalRepository) FindById(id string) (*models.Withdrawal, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	var withdrawal models.Withdrawal
	err = db.Collection("withdrawals").FindOne(ctx, bson.M{"_id": id}).Decode(&withdrawal)
	if err != nil {
		return nil, fmt.Errorf("failed to find withdrawal with id %s: %v", id, err)
	}

	return &withdrawal, nil
}

func (r *withdrawalRepository) FindByUserId(userId string) ([]models.Withdrawal, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	filter := bson.M{"user_id": userId}

	var withdrawals []models.Withdrawal
	cursor, err := db.Collection("withdrawals").Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find withdrawals for user %s: %v", userId, err)
	}

	if err = cursor.All(ctx, &withdrawals); err != nil {
		return nil, fmt.Errorf("failed to read withdrawals cursor: %v", err)
	}

	return withdrawals, nil
}

func (r *withdrawalRepository) InsertOne(withdrawal models.Withdrawal) (*models.Withdrawal, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	withdrawal.WithdrawalId = utils.GenerateRandomID()
	withdrawal.Date = utils.GetCurrentTime()

	_, err = db.Collection("withdrawals").InsertOne(ctx, withdrawal)
	if err != nil {
		return nil, fmt.Errorf("failed to insert withdrawal: %v", err)
	}

	return &withdrawal, nil
}

func (r *withdrawalRepository) DeleteById(id string) error {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	filter := bson.M{"_id": id}
	_, err = db.Collection("withdrawals").DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete withdrawal with id %s: %v", id, err)
	}

	return nil
}

func (r *withdrawalRepository) UpdateById(id string, updatedWithdrawal *models.Withdrawal) (*models.Withdrawal, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedWithdrawal}

	_, err = db.Collection("withdrawals").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update withdrawal with id %s: %v", id, err)
	}

	return updatedWithdrawal, nil
}
