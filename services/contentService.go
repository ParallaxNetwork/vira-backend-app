package services

import (
	"reflect"
	"vira-backend-app/models"
	"vira-backend-app/repositories"
)

type ContentService interface{
	FindAll() ([]models.Content, error)
	FindById(id string) (*models.Content, error)
	InsertOne(content *models.Content) (*models.Content, error)
	DeleteById(id string) error
	UpdateById(id string, updatedContent models.Content) (*models.Content, error)
}

type contentService struct {
	repo repositories.ContentRepository
}

func NewContentService(repo repositories.ContentRepository) ContentService {
	return &contentService{repo: repo}
}

func (s *contentService) FindAll() ([]models.Content, error) {
	return s.repo.FindAll()
}

func (s *contentService) FindById(id string) (*models.Content, error) {
	return s.repo.FindById(id)
}

func (s *contentService) InsertOne(content *models.Content) (*models.Content, error) {
	return s.repo.InsertOne(content)
}

func (s *contentService) DeleteById(id string) error {
	return s.repo.DeleteById(id)
}

func (s *contentService) UpdateById(id string, updatedContent models.Content) (*models.Content, error) {
	content, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	updatedFields := reflect.ValueOf(updatedContent)
	contentValue := reflect.ValueOf(content).Elem()

	for i := 0; i < updatedFields.NumField(); i++ {
		updatedFieldValue := updatedFields.Field(i)
		if !reflect.DeepEqual(updatedFieldValue.Interface(), reflect.Zero(updatedFieldValue.Type()).Interface()) {
			contentValue.Field(i).Set(updatedFieldValue)
		}
	}

	updatedContent_, err := s.repo.UpdateById(id, content)
	if err != nil {
		return nil, err
	}

	return updatedContent_, nil
}
