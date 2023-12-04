package handlers

import (
	"net/http"
	"vira-backend-app/models"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)

type TrackingHandler interface {
	TrackingFindByIdHandler() gin.HandlerFunc
	TrackingFindAllHandler() gin.HandlerFunc
	TrackingInsertHandler() gin.HandlerFunc
	TrackingDeleteByIdHandler() gin.HandlerFunc
	TrackingUpdateByIdHandler() gin.HandlerFunc
}

type trackingHandler struct {
	svc services.TrackingService
}

func NewTrackingHandler(svc services.TrackingService) TrackingHandler {
	return &trackingHandler{
		svc: svc,
	}
}

func (s *trackingHandler) TrackingFindByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		tracking, err := s.svc.FindById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, tracking)
	}
}

func (s *trackingHandler) TrackingFindAllHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		trackings, err := s.svc.FindAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, trackings)
	}
}

func (s *trackingHandler) TrackingInsertHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tracking models.Tracking
		if err := c.ShouldBindJSON(&tracking); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := s.svc.InsertOne(tracking)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func (s *trackingHandler) TrackingDeleteByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := s.svc.DeleteById(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "tracking deleted"})
	}
}

func (s *trackingHandler) TrackingUpdateByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var updatedTracking models.Tracking
		if err := c.ShouldBindJSON(&updatedTracking); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tracking, err := s.svc.UpdateById(id, updatedTracking)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tracking"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Tracking updated successfully", "tracking": tracking})
	}
}
