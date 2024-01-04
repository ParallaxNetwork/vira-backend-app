package routers

import (
	"vira-backend-app/handlers"
	"vira-backend-app/repositories"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)

func WithdrawalRoutes(withdrawal *gin.RouterGroup) {
	withdrawalRepository := repositories.NewWithdrawalRepository()
	withdrawalService := services.NewWithdrawalService(withdrawalRepository)
	withdrawalHandler := handlers.NewWithdrawalHandler(withdrawalService)

	withdrawal.GET("", withdrawalHandler.WithdrawalFindAllHandler())
	withdrawal.GET("/:id", withdrawalHandler.WithdrawalFindByIdHandler())
	withdrawal.GET("/user/:user_id", withdrawalHandler.WithdrawalFindByUserIdHandler())
	withdrawal.POST("", withdrawalHandler.WithdrawalInsertOneHandler())
	withdrawal.PUT("/:id", withdrawalHandler.WithdrawalUpdateByIdHandler())
	withdrawal.DELETE("/:id", withdrawalHandler.WithdrawalDeleteByIdHandler())
}
