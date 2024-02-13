package routers

import (
	"vira-backend-app/handlers"
	"vira-backend-app/repositories"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(admin *gin.RouterGroup) {
	adminRepository := repositories.NewAdminRepository()
	adminService := services.NewAdminService(adminRepository)
	adminHandler := handlers.NewAdminHandler(adminService)

	admin.POST("", adminHandler.AdminInsertHandler())
	admin.GET("", adminHandler.AdminSignInHandler())
}
