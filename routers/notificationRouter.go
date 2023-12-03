package routers

import (
	"vira-backend-app/handlers"
	"vira-backend-app/repositories"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)

func NotificationRoutes(route *gin.RouterGroup) {
	notificationRepository := repositories.NewNotificationRepository()
	notificationService := services.NewNotificationService(notificationRepository)
	notificationHandler := handlers.NewNotificationHandler(notificationService)

	notification := route.Group("/notification")
	{
		notification.GET("", notificationHandler.NotificationPush())
	}
}