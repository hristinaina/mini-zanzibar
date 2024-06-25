package main

import (
	"back/config"
	"back/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.SetTrustedProxies(nil)

	// load data from .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Loading .env file error")
	}

	db, err := config.SetupPostgres()
	if err != nil {
		fmt.Println("could not set up database: %v", err)
	}

	routes.SetupRoutes(router, db)

	// Run the server with HTTPS
	if err := router.RunTLS(":443", "../server.crt", "../server.key"); err != nil {
		fmt.Println("failed to run server: %v", err)
	}
}
