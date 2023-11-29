package services

import (
	"vira-backend-app/models"
	"vira-backend-app/repositories"
)

type TrackingService interface {
	FindById(trackingID string) (*models.Tracking, error)
	FindAll() ([]models.Tracking, error)
	InsertOne(tracking models.Tracking) (*models.Tracking, error)
	DeleteById(trackingId string) (error)
}

type trackingService struct {
	repo repositories.TrackingRepository
}

func NewTrackingService(repo repositories.TrackingRepository) TrackingService {
	return &trackingService{repo: repo}
}

func (s *trackingService) FindById(trackingID string) (*models.Tracking, error) {
	return s.repo.FindById(trackingID)
}

func (s *trackingService) FindAll() ([]models.Tracking, error) {
	return s.repo.FindAll()
}

func (s *trackingService) InsertOne(tracking models.Tracking) (*models.Tracking, error) {
	return s.repo.InsertOne(tracking)
}

func (s *trackingService) DeleteById(trackingId string) (error) {
	return s.repo.DeleteById(trackingId)
}
