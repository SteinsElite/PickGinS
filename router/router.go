package router

import (
	"time"

	_ "github.com/SteinsElite/pickGinS/docs"
	"github.com/SteinsElite/pickGinS/logging"
	"github.com/SteinsElite/pickGinS/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @Title pick finance api
// @Version 0.0.1
// @Description API to interact with the pick finance backend
// @BasePath /api/v1

func SetupGinServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	ginLogger := middleware.GinzapWithConfig(
		logging.Z(),
		&middleware.Config{
			TimeFormat: time.RFC3339,
			UTC:        true,
			SkipPaths: []string{
				"/swagger/index.html",
				"/swagger/swagger-ui-standalone-preset.js",
				"/swagger/swagger-ui.css",
				"/swagger/swagger-ui-bundle.js",
				"/swagger/favicon-32x32.png",
				"/swagger/doc.json",
			},
		},
	)

	r.Use(ginLogger)
	r.Use(middleware.RecoveryWithZap(logging.Z(), false))

	// swagger
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")

	apiv1.GET("/", func(c *gin.Context) {
		c.JSON(200, "pick finance api v1")
	})

	apiv1.GET("/transaction/:address", GetTransaction)

	apiv1.GET("/price_info/:coin", GetCoinPriceInfo)

	apiv1.GET("/notification", GetNotification)
	apiv1.POST("/notification", PublishNotification)

	// api to get the info about vault
	apiv1.GET("/chart/volume", GetVolume)
	apiv1.GET("/chart/profit", GetProfit)
	apiv1.GET("/chart/ratio", GetRatio)

	return r
}
