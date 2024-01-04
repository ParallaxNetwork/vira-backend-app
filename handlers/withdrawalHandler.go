package handlers

import (
	"net/http"
	"vira-backend-app/models"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)

type WithdrawalHandler interface {
	WithdrawalFindAllHandler() gin.HandlerFunc
	WithdrawalFindByIdHandler() gin.HandlerFunc
	WithdrawalFindByUserIdHandler() gin.HandlerFunc
	WithdrawalInsertOneHandler() gin.HandlerFunc
	WithdrawalUpdateByIdHandler() gin.HandlerFunc
	WithdrawalDeleteByIdHandler() gin.HandlerFunc
}

type WithdrawalService struct {
	svc services.WithdrawalService
}

func NewWithdrawalHandler(svc services.WithdrawalService) WithdrawalHandler {
	return &WithdrawalService{
		svc: svc,
	}
}

func (s *WithdrawalService) WithdrawalFindAllHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		withdrawals, err := s.svc.FindAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get withdrawals"})
			return
		}

		c.JSON(http.StatusOK, withdrawals)
	}
}

func (s *WithdrawalService) WithdrawalFindByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		withdrawal, err := s.svc.FindById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Withdrawal not found"})
			return
		}

		c.JSON(http.StatusOK, withdrawal)
	}
}

func (s *WithdrawalService) WithdrawalFindByUserIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		withdrawals, err := s.svc.FindByUserId(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get withdrawals for the user"})
			return
		}

		if len(withdrawals) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No withdrawals found for the user"})
			return
		}

		c.JSON(http.StatusOK, withdrawals)
	}
}

func (s *WithdrawalService) WithdrawalInsertOneHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var withdrawal models.Withdrawal
		if err := c.ShouldBindJSON(&withdrawal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		insertedWithdrawal, err := s.svc.InsertOne(withdrawal)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, insertedWithdrawal)
	}
}

func (s *WithdrawalService) WithdrawalUpdateByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		withdrawalId := c.Param("id")

		var updatedWithdrawal models.Withdrawal
		if err := c.ShouldBindJSON(&updatedWithdrawal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updated, err := s.svc.UpdateById(withdrawalId, &updatedWithdrawal)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update withdrawal"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Withdrawal updated successfully", "withdrawal": updated})
	}
}

func (s *WithdrawalService) WithdrawalDeleteByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		withdrawalId := c.Param("id")

		err := s.svc.DeleteById(withdrawalId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete withdrawal"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Withdrawal deleted successfully"})
	}
}
