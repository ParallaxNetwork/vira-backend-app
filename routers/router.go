package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	mainRouter := gin.Default()

	mainRouter.MaxMultipartMemory = 5 << 20
	mainRouter.Use(CORSMiddleware())
	// mainRouter.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))

	v1 := mainRouter.Group("/api/v1")

	UserRoutes(v1.Group("/user"))
	ProjectRoutes(v1.Group("/project"))
	TrackingRoutes(v1.Group("/tracking"))
	FundingRoutes(v1.Group("/funding"))
	WithdrawalRoutes(v1.Group("/withdrawal"))
	// NotificationRoutes(v1)
	// ContentRoutes(v1)

	return mainRouter
}
