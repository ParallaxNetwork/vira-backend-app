package routers

import (
	"vira-backend-app/handlers"
	"vira-backend-app/repositories"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)

func TrackingRoutes(router *gin.RouterGroup) {
	trackingRepository := repositories.NewTrackingRepository()
	trackingService := services.NewTrackingService(trackingRepository)
	trackingHandler := handlers.NewTrackingHandler(trackingService)

	tracking := router.Group("/tracking")
	{
		tracking.GET("/:id", trackingHandler.TrackingFindByIdHandler())
		tracking.GET("/", trackingHandler.TrackingFindAllHandler())
		tracking.POST("/", trackingHandler.TrackingInsertHandler())
		tracking.DELETE("/:id", trackingHandler.TrackingDeleteByIdHandler())
	}
}
