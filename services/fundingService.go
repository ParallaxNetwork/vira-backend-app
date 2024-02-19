package services

import (
	"reflect"
	"vira-backend-app/dto"
	"vira-backend-app/models"
	"vira-backend-app/repositories"
)

type FundingService interface {
	FindById(id string) (*models.Funding, error)
	FindAll() ([]models.Funding, error)
	FindAllByUserId(userId string) ([]models.Funding, error)
	InsertOne(funding dto.FundingInsertOneDTO) (*models.Funding, error)
	DeleteById(id string) error
	UpdateById(id string, updatedFunding models.Funding) (*models.Funding, error)
}

type fundingService struct {
	repo repositories.FundingRepository
}

func NewFundingService(repo repositories.FundingRepository) FundingService {
	return &fundingService{repo: repo}
}

func (s *fundingService) FindById(id string) (*models.Funding, error) {
	return s.repo.FindById(id)
}

func (s *fundingService) FindAll() ([]models.Funding, error) {
	return s.repo.FindAll()
}

func (s *fundingService) FindAllByUserId(userId string) ([]models.Funding, error) {
	return s.repo.FindAllByUserId(userId)
}

func (s *fundingService) InsertOne(funding dto.FundingInsertOneDTO) (*models.Funding, error) {
	return s.repo.InsertOne(funding)
}

func (s *fundingService) DeleteById(id string) error {
	return s.repo.DeleteById(id)
}

func (s *fundingService) UpdateById(id string, updatedFunding models.Funding) (*models.Funding, error) {
	funding, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	updatedFields := reflect.ValueOf(updatedFunding)
	fundingValue := reflect.ValueOf(funding).Elem()

	for i := 0; i < updatedFields.NumField(); i++ {
		updatedFieldValue := updatedFields.Field(i)
		if !reflect.DeepEqual(updatedFieldValue.Interface(), reflect.Zero(updatedFieldValue.Type()).Interface()) {
			fundingValue.Field(i).Set(updatedFieldValue)
		}
	}

	updatedFunding_, err := s.repo.UpdateById(id, funding)
	if err != nil {
		return nil, err
	}

	return updatedFunding_, nil
}
