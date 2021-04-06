package controller

import (
	"blog/server/model"
	"blog/server/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := service.GetArticleById(id)
	c.JSON(http.StatusOK, gin.H{
		"status": model.Success,
		"data":   data,
	})
}

func PageArticleByCategoryPrefix(c *gin.Context) {
	page := model.Page{}
	if err := c.BindJSON(&page); err != nil {
		panic(err)
	}
	prefix := c.Param("prefix")
	service.PageArticleByCategoryPrefix(prefix, &page)
	c.JSON(http.StatusOK, gin.H{
		"status": model.Success,
		"data":   page,
	})
}
