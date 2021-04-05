package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/service"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	data := service.GetArticleById(&id)

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   data,
	})
}
