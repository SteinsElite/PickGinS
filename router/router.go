package router

import (
	_ "github.com/SteinsElite/pickGinS/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetupGinServer() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiv1 := r.Group("/api/v1")
	apiv1.GET("/", func(c *gin.Context) {
		c.JSON(200, "pick api")
	})
	
	apiv1.GET("/price_info/:coin_ids", GetCoinPriceInfo)
	// the api to get the transaction record
	apiv1.GET("/transaction/:address", GetTransaction)

	//the api to get the global announcement
	apiv1.GET("/announcement/", GetAnn)
	apiv1.GET("/announcement/:category", GetSpecificAnn)

	// api to get the info about vault
	apiv1.GET("/chart/volume", GetVolume)
	apiv1.GET("/chart/profit", GetProfit)
	apiv1.GET("/chart/ratio", GetRatio)

	// the api to get the
	return r
}
