package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}

func KeyNotFoundError(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Key not found",
	})
}
