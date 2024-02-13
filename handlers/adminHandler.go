package handlers

import (
	"net/http"
	"vira-backend-app/models"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)

type AdminHandler interface {
	AdminInsertHandler() gin.HandlerFunc
	AdminSignInHandler() gin.HandlerFunc
}

type adminHandler struct {
	svc services.AdminService
}

func NewAdminHandler(svc services.AdminService) AdminHandler {
	return &adminHandler{
		svc: svc,
	}
}

func (s *adminHandler) AdminInsertHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var admin models.Admin
		if err := c.ShouldBindJSON(&admin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		checkAdmin, _ := s.svc.FindOneByUsername(admin.UserName)
		if checkAdmin == nil {
			checkAdmin, _ = s.svc.FindOneByEmail(admin.Email)
		}

		if checkAdmin != nil {
			c.JSON(http.StatusOK, gin.H{"message": "Admin already exists"})
			return
		}

		insertedAdmin, err := s.svc.InsertOne(admin)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Admin inserted successfully", "admin": insertedAdmin})
	}
}

func (s *adminHandler) AdminSignInHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		usernameEmail := c.Query("usernameEmail")
		password := c.Query("password")

		admin, err := s.svc.SignIn(usernameEmail, password)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, admin)
	}
}
