package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    // Kreira novi Gin router
    router := gin.Default()

    // Definira rutu za početnu stranicu
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, Gin!")
    })

    // Pokreće server na portu 8080
    router.Run(":9000")
}
