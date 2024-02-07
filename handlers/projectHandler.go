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
	ProjectUpdateByIdHandler() gin.HandlerFunc
	StartProject() gin.HandlerFunc
	RevokeProject() gin.HandlerFunc
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

func (s *projectHandler) ProjectUpdateByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId := c.Param("id")

		var updatedProject models.Project
		if err := c.ShouldBindJSON(&updatedProject); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		project, err := s.svc.UpdateById(projectId, updatedProject)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully", "project": project})
	}
}

func (s *projectHandler) StartProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId := c.Param("id")

		err := s.svc.StartProject(projectId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start project"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Project started successfully"})
	}
}

func (s *projectHandler) RevokeProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId := c.Param("id")

		err := s.svc.RevokeProject(projectId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to revoke project"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Project revoked successfully"})
	}
}
