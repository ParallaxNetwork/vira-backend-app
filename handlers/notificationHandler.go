package handlers

import (
	"net/http"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)

type NotificationHandler interface {
	NotificationPush() gin.HandlerFunc
}

type notificationHandler struct {
	svc services.NotificationService
}

func NewNotificationHandler(svc services.NotificationService) NotificationHandler {
	return &notificationHandler{svc: svc}
}

func (s *notificationHandler) NotificationPush() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := s.svc.NotificationPush()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}