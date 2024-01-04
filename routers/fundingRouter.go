package routers

import (
	"vira-backend-app/handlers"
	"vira-backend-app/repositories"
	"vira-backend-app/services"

	"github.com/gin-gonic/gin"
)

func FundingRoutes(funding *gin.RouterGroup) {
	fundingRepository := repositories.NewFundingRepository()
	fundingService := services.NewFundingService(fundingRepository)
	fundingHandler := handlers.NewFundingHandler(fundingService)

	funding.GET("", fundingHandler.FundingFindAllHandler())
	funding.GET("/:id", fundingHandler.FundingFindByIdHandler())
	funding.GET("/user/:id", fundingHandler.FundingFindAllByUserId())
	funding.POST("", fundingHandler.FundingInsertHandler())
	funding.DELETE("/:id", fundingHandler.FundingDeleteByIdHandler())
	funding.PUT("/:id", fundingHandler.FundingUpdateByIdHandler())
}
