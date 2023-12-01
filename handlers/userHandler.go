package handlers

import (
	"net/http"
	"vira-backend-app/models"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	UserFindAllHandler() gin.HandlerFunc
	UserFindByIdHandler() gin.HandlerFunc
	UserInsertHandler() gin.HandlerFunc
	UserDeleteByIdHandler() gin.HandlerFunc
}

type userHandler struct {
	svc services.UserService
}

func NewUserHandler(svc services.UserService) UserHandler {
	return &userHandler{
		svc: svc,
	}
}

func (s *userHandler) UserFindAllHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        users, err := s.svc.FindAll()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
            return
        }

        c.JSON(http.StatusOK, users)
    }
}

func (s *userHandler) UserFindByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")

		user, err := s.svc.FindById(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func (s *userHandler) UserInsertHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		insertedUser, err := s.svc.InsertOne(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User inserted successfully", "user": insertedUser})
	}
}

func (s *userHandler) UserDeleteByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")

		err := s.svc.DeleteById(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}