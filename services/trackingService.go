package services

import (
	"reflect"
	"vira-backend-app/models"
	"vira-backend-app/repositories"
)

type TrackingService interface {
	FindById(trackingID string) (*models.Tracking, error)
	FindAll() ([]models.Tracking, error)
	InsertOne(tracking models.Tracking) (*models.Tracking, error)
	DeleteById(trackingId string) (error)
	UpdateById(trackingID string, updatedTracking models.Tracking) (*models.Tracking, error)
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

func (s *trackingService) UpdateById(trackingID string, updatedTracking models.Tracking) (*models.Tracking, error) {
  tracking, err := s.repo.FindById(trackingID)
	if err != nil {
		return nil, err
	}

	updatedValue := reflect.ValueOf(updatedTracking)
	trackingValue := reflect.ValueOf(tracking).Elem()

	for i := 0; i < updatedValue.NumField(); i++ {
		updatedFieldValue := updatedValue.Field(i)
		if !reflect.DeepEqual(updatedFieldValue.Interface(), reflect.Zero(updatedFieldValue.Type()).Interface()) {
			trackingValue.Field(i).Set(updatedFieldValue)
		}
	}

	updatedTracking_, err := s.repo.UpdateById(trackingID, tracking)
	if err != nil {
		return nil, err
	}

	return updatedTracking_, nil
}
