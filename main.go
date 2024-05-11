package main

import (
	"be-assignment/internal/config"
	"be-assignment/internal/database"
	accountmanager "be-assignment/internal/modules/account-manager"
	"log"

	docs "be-assignment/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title ConcreteAI-Assignment API
// @version 1.0
// @description This is a simple API for Banking Core System.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	conf := config.NewConfig()

	db := database.GetDatabaseConnection(conf)

	database.Migrate(db)

	userRepository := accountmanager.NewUserRepository(db)
	accountRepository := accountmanager.NewAccountRepository(db)

	accountManagerService := accountmanager.NewService(userRepository, accountRepository)

	app := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	accountmanager.NewRoute(app, accountManagerService)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Fatal(app.Run(":" + conf.Srv.Port))
}
