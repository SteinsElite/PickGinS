package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/SteinsElite/pickGinS/docs"
)

// @Title pick finance api
// @Version 0.0.1
// @Description API to interact with the pick finance backend
// @BasePath /api/v1

func SetupGinServer() *gin.Engine {
	r := gin.New()

	// middleware of logger and recovery of gin default now
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// swagger
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")

	apiv1.GET("/", func(c *gin.Context) {
		c.JSON(200, "pick finance api v1")
	})
	
	apiv1.GET("/transaction/:address", GetTransaction)
	apiv1.GET("/price_info/:coin", GetCoinPriceInfo)


	apiv1.GET("/notification", GetNotification)

	// api to get the info about vault
	apiv1.GET("/chart/volume", GetVolume)
	apiv1.GET("/chart/profit", GetProfit)
	apiv1.GET("/chart/ratio", GetRatio)

	apiv1.GET("/auth/keyword_hash", GetKeyWordHash)
	apiv1.POST("/auth/:address/add_publisher", AddPublisher)
	apiv1.POST("/notification/:publisher", PublishNotification)

	return r
}
