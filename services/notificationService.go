package services

import (
	"vira-backend-app/repositories"
)

type NotificationService interface {
	NotificationPush() (error)
}

type notificationService struct {
	repo repositories.NotificationRepository
}

func NewNotificationService(repo repositories.NotificationRepository) NotificationService {
	return &notificationService{repo: repo}
}

func (s *notificationService) NotificationPush() (error) {
	return s.repo.NotificationPush()
}