package services

import (
	"reflect"
	"vira-backend-app/models"
	"vira-backend-app/repositories"
)

type ProjectService interface {
	FindById(projectID string) (*models.Project, error)
	FindAll() ([]models.Project, error)
	InsertOne(project models.Project) (*models.Project, error)
	DeleteById(projectId string) (error)
	UpdateById(projectID string, updatedProject models.Project) (*models.Project, error)
}

type projectService struct {
	repo repositories.ProjectRepository
}

func NewProjectService(repo repositories.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

func (s *projectService) FindById(projectID string) (*models.Project, error) {
	return s.repo.FindById(projectID)
}

func (s *projectService) FindAll() ([]models.Project, error) {
	return s.repo.FindAll()
}

func (s *projectService) InsertOne(project models.Project) (*models.Project, error) {
	return s.repo.InsertOne(project)
}

func (s *projectService) DeleteById(projectId string) (error) {
	return s.repo.DeleteById(projectId)
}

func (s *projectService) UpdateById(projectID string, updatedProject models.Project) (*models.Project, error) {
	project, err := s.repo.FindById(projectID)
	if err != nil {
		return nil, err
	}

	updatedFields := reflect.ValueOf(updatedProject)
	projectValue := reflect.ValueOf(project).Elem()

	for i := 0; i < updatedFields.NumField(); i++ {
		updatedFieldValue := updatedFields.Field(i)
		if !reflect.DeepEqual(updatedFieldValue.Interface(), reflect.Zero(updatedFieldValue.Type()).Interface()) {
			projectValue.Field(i).Set(updatedFieldValue)
		}
	}

	updatedProject_, err := s.repo.UpdateById(projectID, project)
	if err != nil {
		return nil, err
	}

	return updatedProject_, nil
}
