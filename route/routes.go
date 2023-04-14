package route

import (
	controllers "sellerApp/src/controllers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine) {

	/***BASEPATH OF AN API. NOTE:THIS SHOULDN'T BE CHANGED***/
	api := router.Group("/api/v1")

	/***API ROUTES***/
	api.POST("/placeOrder", controllers.PlaceOrder)
	api.POST("/cancelOrder", controllers.CancelOrder)
	router.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(viper.GetString("server.port"))
}
