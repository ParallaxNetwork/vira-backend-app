package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	mainRouter := gin.Default()

	UserRoutes(mainRouter)
	ProjectRoutes(mainRouter)
	FundingRoutes(mainRouter)
	TrackingRoutes(mainRouter)

	return mainRouter
}
