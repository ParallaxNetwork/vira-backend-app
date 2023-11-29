package handlers

import (
	"net/http"
	"vira-backend-app/models"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)

type ProjectHandler interface {
	ProjectFindByIdHandler() gin.HandlerFunc
	ProjectFindAllHandler() gin.HandlerFunc
	ProjectInsertHandler() gin.HandlerFunc
	ProjectDeleteByIdHandler() gin.HandlerFunc
}

type projectHandler struct {
	svc services.ProjectService
}

func NewProjectHandler(svc services.ProjectService) ProjectHandler {
	return &projectHandler{svc: svc}
}

func (s *projectHandler) ProjectFindByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId := c.Param("id")

		project, err := s.svc.FindById(projectId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get project"})
			return
		}

		c.JSON(http.StatusOK, project)
	}
}

func (s *projectHandler) ProjectFindAllHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		projects, err := s.svc.FindAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get projects"})
			return
		}

		c.JSON(http.StatusOK, projects)
	}
}

func (s *projectHandler) ProjectInsertHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var project models.Project
		if err := c.ShouldBindJSON(&project); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		insertedProject, err := s.svc.InsertOne(project)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Project inserted successfully", "project": insertedProject})
	}
}

func (s *projectHandler) ProjectDeleteByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId := c.Param("id")

		err := s.svc.DeleteById(projectId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
	}
}