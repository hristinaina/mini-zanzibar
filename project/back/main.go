package main

import (
	"back/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load data from .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Loading .env file error")
	}

	// Kreira novi Gin router
	router := gin.Default()

	config.SetupPostgres()

	// Definira rutu za početnu stranicu
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")
	})

	// Pokreće server na portu 8080
	router.Run(":9000")
}
