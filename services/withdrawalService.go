package services

import (
	"reflect"
	"vira-backend-app/models"
	"vira-backend-app/repositories"
)

type WithdrawalService interface {
	FindAll() ([]models.Withdrawal, error)
	FindById(id string) (*models.Withdrawal, error)
	FindByUserId(userId string) ([]models.Withdrawal, error)
	InsertOne(withdrawal models.Withdrawal) (*models.Withdrawal, error)
	UpdateById(id string, updatedWithdrawal *models.Withdrawal) (*models.Withdrawal, error)
	DeleteById(id string) error
}

type withdrawalService struct {
	repo repositories.WithdrawalRepository
}

func NewWithdrawalService(repo repositories.WithdrawalRepository) WithdrawalService {
	return &withdrawalService{repo: repo}
}

func (s *withdrawalService) FindAll() ([]models.Withdrawal, error) {
	return s.repo.FindAll()
}

func (s *withdrawalService) FindById(id string) (*models.Withdrawal, error) {
	return s.repo.FindById(id)
}

func (s *withdrawalService) FindByUserId(userId string) ([]models.Withdrawal, error) {
	return s.repo.FindByUserId(userId)
}

func (s *withdrawalService) InsertOne(withdrawal models.Withdrawal) (*models.Withdrawal, error) {
	return s.repo.InsertOne(withdrawal)
}

func (s *withdrawalService) UpdateById(id string, updatedWithdrawal *models.Withdrawal) (*models.Withdrawal, error) {
	withdrawal, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	updatedFields := reflect.ValueOf(*updatedWithdrawal)
	withdrawalValue := reflect.ValueOf(withdrawal).Elem()

	for i := 0; i < updatedFields.NumField(); i++ {
		updatedFieldValue := updatedFields.Field(i)
		if !reflect.DeepEqual(updatedFieldValue.Interface(), reflect.Zero(updatedFieldValue.Type()).Interface()) {
			withdrawalValue.Field(i).Set(updatedFieldValue)
		}
	}

	updatedWithdrawal_, err := s.repo.UpdateById(id, withdrawal)
	if err != nil {
		return nil, err
	}

	return updatedWithdrawal_, nil
}

func (s *withdrawalService) DeleteById(id string) error {
	return s.repo.DeleteById(id)
}