package services

import (
	"vira-backend-app/models"
	"vira-backend-app/repositories"
)

type AdminService interface {
	InsertOne(admin models.Admin) (*models.Admin, error)
	FindOneByUsername(username string) (*models.Admin, error)
	FindOneByEmail(email string) (*models.Admin, error)
	SignIn(usernameEmail string, password string) (*models.Admin, error)
}

type adminService struct {
	repo repositories.AdminRepository
}

func NewAdminService(repo repositories.AdminRepository) AdminService {
	return &adminService{repo: repo}
}

func (s *adminService) InsertOne(admin models.Admin) (*models.Admin, error) {
	return s.repo.InsertOne(admin)
}

func (s *adminService) SignIn(usernameEmail string, password string) (*models.Admin, error) {
	return s.repo.SignIn(usernameEmail, password)
}

func (s *adminService) FindOneByUsername(username string) (*models.Admin, error) {
	return s.repo.FindOneByUsername(username)
}

func (s *adminService) FindOneByEmail(email string) (*models.Admin, error) {
	return s.repo.FindOneByEmail(email)
}
