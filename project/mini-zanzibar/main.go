package main

import (
	"fmt"
	"mini-zanzibar/config"
	"mini-zanzibar/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	router.Use(config.SetupCORS())

	// load data from .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Loading .env file error")
	}

	levelDb := config.InitLevelDB()
	defer config.CloseLevelDB(levelDb)

	consulDB := config.InitConsulDB()

	routes.SetupRoutes(router, levelDb, consulDB)
	err := router.Run(":8081")
	if err != nil {
		return
	} // Listen and serve on 0.0.0.0:8081
}
