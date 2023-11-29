package services

import (
	"vira-backend-app/models"
	"vira-backend-app/repositories"
)

type FundingService interface {
	FindById(fundingId string) (*models.Funding, error)
	FindAll() ([]models.Funding, error)
	FindAllByUserId(userId string) ([]models.Funding, error)
	InsertOne(funding models.Funding) (*models.Funding, error)
	DeleteById(fundingId string) (error)
}

type fundingService struct {
	repo repositories.FundingRepository
}

func NewFundingService (repo repositories.FundingRepository) FundingService {
	return &fundingService{repo: repo}
}

func (s *fundingService) FindById(fundingId string) (*models.Funding, error) {
	return s.repo.FindById(fundingId)
}

func (s *fundingService) FindAll() ([]models.Funding, error) {
	return s.repo.FindAll()
}

func (s *fundingService) FindAllByUserId(userId string) ([]models.Funding, error) {
	return s.repo.FindAllByUserId(userId)
}

func (s *fundingService) InsertOne(funding models.Funding) (*models.Funding, error) {
	return s.repo.InsertOne(funding)
}

func (s *fundingService) DeleteById(fundingId string) (error) {
	return s.repo.DeleteById(fundingId)
}