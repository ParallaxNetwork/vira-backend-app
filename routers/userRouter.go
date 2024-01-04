package routers

import (
	"vira-backend-app/handlers"
	"vira-backend-app/repositories"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(user *gin.RouterGroup) {
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	user.GET("/:id", userHandler.UserFindByIdHandler())
	user.GET("", userHandler.UserFindAllHandler())
	user.POST("", userHandler.UserInsertHandler())
	user.DELETE("/:id", userHandler.UserDeleteByIdHandler())
	user.PUT("/:id", userHandler.UserUpdateByIdHandler())
}
