package main

import (
	"be-assignment/internal/config"
	"be-assignment/internal/database"
	"be-assignment/internal/modules/user"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.NewConfig()

	db := database.GetDatabaseConnection(conf)

	database.Migrate(db)

	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository)

	app := gin.Default()

	user.NewRoute(app, userService)

	log.Fatal(app.Run(":" + conf.Srv.Port))
}
