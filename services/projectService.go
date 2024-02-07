package services

import (
	"reflect"
	"vira-backend-app/models"
	"vira-backend-app/repositories"
)

type ProjectService interface {
	FindById(id string) (*models.Project, error)
	FindAll() ([]models.Project, error)
	InsertOne(project models.Project) (*models.Project, error)
	DeleteById(id string) error
	UpdateById(id string, updatedProject models.Project) (*models.Project, error)
	StartProject(projectId string) error
	RevokeProject(projectId string) error
}

type projectService struct {
	repo repositories.ProjectRepository
}

func NewProjectService(repo repositories.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

func (s *projectService) FindById(id string) (*models.Project, error) {
	return s.repo.FindById(id)
}

func (s *projectService) FindAll() ([]models.Project, error) {
	return s.repo.FindAll()
}

func (s *projectService) InsertOne(project models.Project) (*models.Project, error) {
	return s.repo.InsertOne(project)
}

func (s *projectService) DeleteById(id string) error {
	return s.repo.DeleteById(id)
}

func (s *projectService) UpdateById(id string, updatedProject models.Project) (*models.Project, error) {
	project, err := s.repo.FindById(id)
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

	updatedProject_, err := s.repo.UpdateById(id, project)
	if err != nil {
		return nil, err
	}

	return updatedProject_, nil
}

func (s *projectService) StartProject(projectId string) error {
	return s.repo.StartProject(projectId)
}

func (s *projectService) RevokeProject(projectId string) error {
	return s.repo.RevokeProject(projectId)
}
