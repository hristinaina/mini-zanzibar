package main

import (
	"back/config"
	"back/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	// load data from .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Loading .env file error")
	}

	db, _ := config.SetupPostgres()

	routes.SetupRoutes(router, db)

	router.Run(":9000")
}
