package routers

import (
	"vira-backend-app/handlers"
	"vira-backend-app/repositories"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.RouterGroup) {
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	user := route.Group("/user")
	{
		user.GET("", userHandler.UserFindAllHandler())
		user.GET("/:id", userHandler.UserFindByIdHandler())
		user.POST("", userHandler.UserInsertHandler())
		user.DELETE("/:id", userHandler.UserDeleteByIdHandler())
	}
}