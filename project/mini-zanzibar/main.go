package main

import (
	"fmt"
	"mini-zanzibar/config"
	"mini-zanzibar/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.SetTrustedProxies(nil)
	router.Use(config.SetupCORS())

	// load data from .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Loading .env file error")
	}

	levelDb := config.InitLevelDB()
	defer config.CloseLevelDB(levelDb)

	consulDB := config.InitConsulDB()

	routes.SetupRoutes(router, levelDb, consulDB)
	// Run the server with HTTPS
	if err := router.RunTLS(":8443", "../server.crt", "../server.key"); err != nil {
		fmt.Println("failed to run server: %v", err)
	}
}
