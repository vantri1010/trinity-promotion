package main

import (
	"trinity/boots"
	"trinity/config"
	"trinity/modules/router"
	"trinity/utils/logger"
)

var log logger.Logger

// @title Trinity App API
// @version 1.0
// @description API documentation for Trinity App

// @contact.name API Support
// @contact.url http://www.vantri1010.com/support
// @contact.email vantri1010@gmail.com

// @host localhost:8080
// @BasePath /
func main() {
	log = logger.NewLogger("main")
	// Load environment variables
	env := config.NewEnv()

	app, err := boots.InitApp(env)
	if err != nil {
		log.Fatalf("Failed to boots application: %v", err)
	}
	log.Info("Application initialized")

	// Set up the router
	r := router.SetupRouter(app)
	log.Info("Router set up")

	// Start the server
	log.Infof("Server is running on port %s", 8080)
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
