package main

import (
	"github.com/gin-gonic/gin"
	"trinity/bootstrap"
	"trinity/route"
)

// @title Trinity App API
// @version 1.0
// @description API documentation for Trinity App

// @contact.name API Support
// @contact.url http://www.vantri1010.com/support
// @contact.email vantri1010@gmail.com

// @host localhost:8080
// @BasePath /
func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	gin := gin.Default()

	route.SetupRouter(env, db, gin)

	err := gin.Run(env.ServerAddress)
	if err != nil {
		return
	}

}
