package repositories

import (
	"fmt"
	"vira-backend-app/models"
	"vira-backend-app/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type AdminRepository interface {
	InsertOne(admin models.Admin) (*models.Admin, error)
	FindOneByUsername(username string) (*models.Admin, error)
	FindOneByEmail(email string) (*models.Admin, error)
	SignIn(usernameEmail string, password string) (*models.Admin, error)
}

type adminRepository struct{}

func NewAdminRepository() AdminRepository {
	return &adminRepository{}
}

func (r *adminRepository) InsertOne(admin models.Admin) (*models.Admin, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	admin.AdminId = utils.GenerateRandomID()
	admin.CreatedAt = utils.GetCurrentTime()
	admin.Password, _ = utils.HashPassword(admin.Password)

	_, err = db.Collection("admins").InsertOne(ctx, admin)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &admin, nil
}

func (r *adminRepository) FindOneByUsername(username string) (*models.Admin, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var admin models.Admin
	err = db.Collection("admins").FindOne(ctx, bson.M{"username": username}).Decode(&admin)
	if err != nil {
		return nil, fmt.Errorf("username not found")
	}

	return &admin, nil
}

func (r *adminRepository) FindOneByEmail(email string) (*models.Admin, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var admin models.Admin
	err = db.Collection("admins").FindOne(ctx, bson.M{"email": email}).Decode(&admin)
	if err != nil {
		return nil, fmt.Errorf("email not found")
	}

	return &admin, nil
}

func (r *adminRepository) SignIn(usernameEmail string, password string) (*models.Admin, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var admin models.Admin

	db.Collection("admins").FindOne(ctx, bson.M{"username": usernameEmail}).Decode(&admin)
	if admin.AdminId == "" {
		err = db.Collection("admins").FindOne(ctx, bson.M{"email": usernameEmail}).Decode(&admin)
		if err != nil {
			return nil, fmt.Errorf("username/email not found")
		}
	}

	if !utils.CheckPasswordHash(password, admin.Password) {
		return nil, fmt.Errorf("wrong password")
	}

	return &admin, nil
}
