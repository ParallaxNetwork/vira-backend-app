package handlers

import (
	"net/http"
	"vira-backend-app/dto"
	"vira-backend-app/models"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)

type FundingHandler interface {
	FundingFindByIdHandler() gin.HandlerFunc
	FundingFindAllHandler() gin.HandlerFunc
	FundingFindAllByUserId() gin.HandlerFunc
	FundingInsertHandler() gin.HandlerFunc
	FundingDeleteByIdHandler() gin.HandlerFunc
	FundingUpdateByIdHandler() gin.HandlerFunc
}

type fundingHandler struct {
	svc services.FundingService
}

func NewFundingHandler(svc services.FundingService) FundingHandler {
	return &fundingHandler{svc: svc}
}

func (s *fundingHandler) FundingFindByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fundingId := c.Param("id")

		funding, err := s.svc.FindById(fundingId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get funding"})
			return
		}

		c.JSON(http.StatusOK, funding)
	}
}

func (s *fundingHandler) FundingFindAllHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fundings, err := s.svc.FindAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get fundings"})
			return
		}

		c.JSON(http.StatusOK, fundings)
	}
}

func (s *fundingHandler) FundingFindAllByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("id")

		fundings, err := s.svc.FindAllByUserId(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get fundings"})
			return
		}

		c.JSON(http.StatusOK, fundings)
	}
}

func (s *fundingHandler) FundingInsertHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var funding dto.FundingInsertOneDTO
		if err := c.ShouldBindJSON(&funding); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		insertedFunding, err := s.svc.InsertOne(funding)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Funding inserted successfully", "funding": insertedFunding})
	}
}

func (s *fundingHandler) FundingDeleteByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fundingId := c.Param("id")

		err := s.svc.DeleteById(fundingId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete funding"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Funding deleted successfully"})
	}
}

func (s *fundingHandler) FundingUpdateByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fundingId := c.Param("id")

		var updatedFunding models.Funding
		if err := c.ShouldBindJSON(&updatedFunding); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		funding, err := s.svc.UpdateById(fundingId, updatedFunding)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update funding"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Funding updated successfully", "funding": funding})
	}
}
