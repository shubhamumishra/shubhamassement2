package main

import (
	"fmt"

	config "sellerApp/config"
	"sellerApp/docs"
	"sellerApp/migration"
	"sellerApp/route"
	"sellerApp/src/repository"
	"sellerApp/utils/database"
	logger "sellerApp/utils/logging"
	"sellerApp/utils/validator"

	"github.com/spf13/viper"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("App is Ready")
	config.LoadConfig()
	router := gin.Default()
	logger.NewLogger("./log/debug.log")
	database.GetInstancemysql()

	repository.MySqlInit()
	migration.Migration()
	validator.Init()

	docs.SwaggerInfo.Title = "Seller App API"
	docs.SwaggerInfo.Description = "Documentation seller API v1.0"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = viper.Get("swagger.url").(string)
	docs.SwaggerInfo.Schemes = []string{viper.GetString("swagger.method")}
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("authorization")
	router.Use(cors.New(corsConfig))
	route.SetupRoutes(router)

}
