package services

import (
	"reflect"
	"vira-backend-app/models"
	"vira-backend-app/repositories"
)

type UserService interface {
	FindAll() ([]*models.User, error)
	FindById(userID string) (*models.User, error)
	InsertOne(user models.User) (*models.User, error)
	DeleteById(userId string) (error)
	UpdateById(userID string, updatedUser models.User) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) FindAll() ([]*models.User, error) {
  return s.repo.FindAll()
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

func (s *userService) UpdateById(userID string, updatedUser models.User) (*models.User, error) {
	user, err := s.repo.FindById(userID)
	if err != nil {
		return nil, err
	}

	updatedFields := reflect.ValueOf(updatedUser)
	userValue := reflect.ValueOf(user).Elem()

	for i := 0; i < updatedFields.NumField(); i++ {
		updatedFieldValue := updatedFields.Field(i)
		if !reflect.DeepEqual(updatedFieldValue.Interface(), reflect.Zero(updatedFieldValue.Type()).Interface()) {
			userValue.Field(i).Set(updatedFieldValue)
		}
	}

	updatedUser_, err := s.repo.UpdateById(userID, user)
	if err != nil {
		return nil, err
	}

	return updatedUser_, nil
}
