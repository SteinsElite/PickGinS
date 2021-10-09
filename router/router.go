package router

import (
	_ "github.com/SteinsElite/pickGinS/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title pick finance api
// @version 1.0
// @description API to get data from blockchain
// @BasePath /api/v1


func SetupGinServer() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// swagger
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	apiv1.GET("/", func(c *gin.Context) {
		c.JSON(200, "pick api v1")
	})
	
	apiv1.GET("/price_info/:coin", GetCoinPriceInfo)
	// the api to get the transaction record
	apiv1.GET("/transaction/:address", GetTransaction)

	apiv1.GET("/notification", GetNotification)

	// api to get the info about vault
	apiv1.GET("/chart/volume", GetVolume)
	apiv1.GET("/chart/profit", GetProfit)
	apiv1.GET("/chart/ratio", GetRatio)


	// api to manage the notification
	// no need for sig
	apiv1.GET("/auth/word_hash", GetWordHash)
	// need sig to call
	apiv1.POST("/auth/:address/register", RegisterPublisher)


	apiv1.POST("/notification", PublishNotification)

	return r
}
