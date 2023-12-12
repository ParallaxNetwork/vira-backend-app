package services

import (
	"reflect"
	"vira-backend-app/models"
	"vira-backend-app/repositories"
)

type TrackingService interface {
	FindById(id string) (*models.Tracking, error)
	FindAll() ([]models.Tracking, error)
	InsertOne(tracking models.Tracking) (*models.Tracking, error)
	DeleteById(id string) (error)
	UpdateById(id string, updatedTracking models.Tracking) (*models.Tracking, error)
}

type trackingService struct {
	repo repositories.TrackingRepository
}

func NewTrackingService(repo repositories.TrackingRepository) TrackingService {
	return &trackingService{repo: repo}
}

func (s *trackingService) FindById(id string) (*models.Tracking, error) {
	return s.repo.FindById(id)
}

func (s *trackingService) FindAll() ([]models.Tracking, error) {
	return s.repo.FindAll()
}

func (s *trackingService) InsertOne(tracking models.Tracking) (*models.Tracking, error) {
	return s.repo.InsertOne(tracking)
}

func (s *trackingService) DeleteById(id string) (error) {
	return s.repo.DeleteById(id)
}

func (s *trackingService) UpdateById(id string, updatedTracking models.Tracking) (*models.Tracking, error) {
  tracking, err := s.repo.FindById(id)
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

	updatedTracking_, err := s.repo.UpdateById(id, tracking)
	if err != nil {
		return nil, err
	}

	return updatedTracking_, nil
}
