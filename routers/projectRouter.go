package routers

import (
	"vira-backend-app/handlers"
	"vira-backend-app/repositories"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)

func ProjectRoutes(router *gin.Engine) {
	projectRepository := repositories.NewProjectRepository()
	projectService := services.NewProjectService(projectRepository)
	projectHandler := handlers.NewProjectHandler(projectService)

	project := router.Group("/project")
	{
		project.GET("/:id", projectHandler.ProjectFindByIdHandler())
		project.GET("/", projectHandler.ProjectFindAllHandler())
		project.POST("/", projectHandler.ProjectInsertHandler())
		project.DELETE("/:id", projectHandler.ProjectDeleteByIdHandler())
	}
}
