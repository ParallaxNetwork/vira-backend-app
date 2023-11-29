package services

import (
	"vira-backend-app/models"
	"vira-backend-app/repositories"
)

type UserService interface {
	FindById(userID string) (*models.User, error)
	InsertOne(user models.User) (*models.User, error)
	DeleteById(userId string) (error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) FindById(userID string) (*models.User, error) {
	return s.repo.FindById(userID)
}

func (s *userService) InsertOne(user models.User) (*models.User, error) {
	return s.repo.InsertOne(user)
}

func (s *userService) DeleteById(userId string) (error) {
	return s.repo.DeleteById(userId)
}