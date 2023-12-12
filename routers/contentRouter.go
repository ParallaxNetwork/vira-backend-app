package routers

import (
	"vira-backend-app/handlers"
	"vira-backend-app/repositories"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)

func ContentRoutes(route *gin.RouterGroup) {
	contentRepository := repositories.NewContentRepository()
	contentService := services.NewContentService(contentRepository)
	contentHandler := handlers.NewContentHandler(contentService)

	content := route.Group("/content")
	{
		content.GET("", contentHandler.ContentFindAllHandler())
		content.GET("/:id", contentHandler.ContentFindByIdHandler())
		content.POST("", contentHandler.ContentInsertHandler())
		content.DELETE("/:id", contentHandler.ContentDeleteByIdHandler())
		content.PUT("/:id", contentHandler.ContentUpdateByIdHandler())
	}
}