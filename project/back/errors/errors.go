package errors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}

func NotFoundError(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Not found",
	})
}

func BadRequestError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func ReturnError(c *gin.Context, err error) {
	var customErr CustomError
	if errors.As(err, &customErr) {
		switch customErr.Code {
		case 404:
			NotFoundError(c)
		case 500:
			InternalServerError(c, err)
		default:
			BadRequestError(c, err)
		}
	} else {
		BadRequestError(c, err)
	}
}
