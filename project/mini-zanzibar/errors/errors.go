package errors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomError struct {
	Code    int
	Message string
}

// Error implements the error interface for CustomError.
func (e CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

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

func BadRequestError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
