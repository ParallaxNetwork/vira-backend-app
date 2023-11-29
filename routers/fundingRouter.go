package routers

import (
	"vira-backend-app/handlers"
	"vira-backend-app/repositories"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)


func FundingRoutes(router *gin.Engine) {
	fundingRepo := repositories.NewFundingRepository()
	fundingSvc := services.NewFundingService(fundingRepo)
	fundingHandler := handlers.NewFundingHandler(fundingSvc)

	fundingRoutes := router.Group("/funding")
	{
		fundingRoutes.GET("/", fundingHandler.FundingFindAllHandler())
		fundingRoutes.GET("/:id", fundingHandler.FundingFindByIdHandler())
		fundingRoutes.GET("/user/:id", fundingHandler.FundingFindAllByUserId())
		fundingRoutes.POST("/", fundingHandler.FundingInsertHandler())
		fundingRoutes.DELETE("/:id", fundingHandler.FundingDeleteByIdHandler())
	}
}
