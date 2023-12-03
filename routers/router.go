package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	mainRouter := gin.Default()

	v1 := mainRouter.Group("api/v1")

	UserRoutes(v1)
	ProjectRoutes(v1)
	FundingRoutes(v1)
	TrackingRoutes(v1)
	NotificationRoutes(v1)

	return mainRouter
}
