package handlers

import (
	"net/http"
	"vira-backend-app/models"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)

type ContentHandler interface {
	ContentFindAllHandler() gin.HandlerFunc
	ContentFindByIdHandler() gin.HandlerFunc
	ContentInsertHandler() gin.HandlerFunc
	ContentDeleteByIdHandler() gin.HandlerFunc
	ContentUpdateByIdHandler() gin.HandlerFunc
}

type contentHandler struct {
	svc services.ContentService
}

func NewContentHandler(svc services.ContentService) ContentHandler {
	return &contentHandler{svc: svc}
}

func (h *contentHandler) ContentFindAllHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		contents, err := h.svc.FindAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, contents)
	}
}

func (h *contentHandler) ContentFindByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		content, err := h.svc.FindById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, content)
	}
}

func (h *contentHandler) ContentInsertHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var content models.Content
		if err := c.ShouldBindJSON(&content); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		insertedContent, err := h.svc.InsertOne(&content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, insertedContent)
	}
}

func (h *contentHandler) ContentDeleteByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := h.svc.DeleteById(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Content deleted successfully"})
	}
}

func (h *contentHandler) ContentUpdateByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var updatedContent models.Content
		if err := c.ShouldBindJSON(&updatedContent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		content, err := h.svc.UpdateById(id, updatedContent)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, content)
	}
}
