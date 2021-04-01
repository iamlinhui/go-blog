package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    "a",
		"id":      12,
		"message": "msg",
		"token":   "sa",
	})
}

