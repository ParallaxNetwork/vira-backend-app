package services

import (
	"reflect"
	"vira-backend-app/models"
	"vira-backend-app/repositories"
)

type UserService interface {
	FindAll() ([]*models.User, error)
	FindById(id string) (*models.User, error)
	FindByWallet(wallet string) (*models.User, error)
	InsertOne(user models.User) (*models.User, error)
	DeleteById(id string) (error)
	UpdateById(id string, updatedUser models.User) (*models.User, error)
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

func (s *userService) FindById(id string) (*models.User, error) {
	return s.repo.FindById(id)
}

func (s *userService) FindByWallet(wallet string) (*models.User, error) {
	return s.repo.FindByWallet(wallet)
}

func (s *userService) InsertOne(user models.User) (*models.User, error) {
	return s.repo.InsertOne(user)
}

func (s *userService) DeleteById(id string) (error) {
	return s.repo.DeleteById(id)
}

func (s *userService) UpdateById(id string, updatedUser models.User) (*models.User, error) {
	user, err := s.repo.FindById(id)
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

	updatedUser_, err := s.repo.UpdateById(id, user)
	if err != nil {
		return nil, err
	}

	return updatedUser_, nil
}
