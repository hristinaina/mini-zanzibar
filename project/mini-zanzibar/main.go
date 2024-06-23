package main

import (
	"github.com/gin-gonic/gin"
	"mini-zanzibar/config"
	"mini-zanzibar/routes"
)

func main() {
	router := gin.Default()
	router.Use(config.SetupCORS())

	levelDb := config.InitLevelDB()
	defer config.CloseLevelDB(levelDb)

	consulDB := config.InitConsulDB()

	routes.SetupRoutes(router, levelDb, consulDB)
	err := router.Run(":8081")
	if err != nil {
		return
	} // Listen and serve on 0.0.0.0:8081
}
