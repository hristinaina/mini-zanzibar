package main

import (
	"back/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
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
