package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	err := router.Run(":8081")
	if err != nil {
		return
	} // Listen and serve on 0.0.0.0:8081
}
