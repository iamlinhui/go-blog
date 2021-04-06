package controller

import (
	"blog/server/model"
	"blog/server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPublishArticleByUrl(c *gin.Context) {
	data := service.GetPublishArticleByUrl(c.Param("url"))
	c.JSON(http.StatusOK, gin.H{
		"status": model.Success,
		"data":   data,
	})
}

func PagePublishArticle(c *gin.Context) {
	page := model.Page{}
	if err := c.BindJSON(&page); err != nil {
		panic(err)
	}
	service.PagePublishArticle(&page)
	c.JSON(http.StatusOK, gin.H{
		"status": model.Success,
		"data":   page,
	})
}

func PagePublishArticleByCategoryPrefix(c *gin.Context) {
	page := model.Page{}
	if err := c.BindJSON(&page); err != nil {
		panic(err)
	}
	prefix := c.Param("prefix")
	service.PagePublishArticleByCategoryPrefix(prefix, &page)
	c.JSON(http.StatusOK, gin.H{
		"status": model.Success,
		"data":   page,
	})
}
