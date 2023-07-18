package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MessageBadRequest(c *gin.Context, message string) {
	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"status": "error",
		"message": message,
		"data": nil,
	})
}

func MessageUnauthorized(c *gin.Context, message string) {
	c.IndentedJSON(http.StatusUnauthorized, gin.H{
		"status": "error",
		"message": message,
		"data": nil,
	})
}

func MessageInternalError(c *gin.Context, message string) {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{
		"status": "error",
		"message": message,
		"data": nil,
	})
}
