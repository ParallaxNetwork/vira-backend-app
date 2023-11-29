package services

import (
	"vira-backend-app/models"
	"vira-backend-app/repositories"
)

type ProjectService interface {
	FindById(projectID string) (*models.Project, error)
	FindAll() ([]models.Project, error)
	InsertOne(project models.Project) (*models.Project, error)
	DeleteById(projectId string) (error)
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