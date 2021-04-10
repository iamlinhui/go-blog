package controller

import (
	"blog/server/model"
	"blog/server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPublishArticleByUrl(c *gin.Context) {
	data := service.GetPublishArticleByUrl(c.Param("url"))
	c.JSON(http.StatusOK, model.Ok(data))
}

func PagePublishArticle(c *gin.Context) {
	page := model.Page{}
	if err := c.BindJSON(&page); err != nil {
		panic(err)
	}
	service.PagePublishArticle(&page)
	c.JSON(http.StatusOK, model.Ok(page))
}

func PagePublishArticleByCategoryPrefix(c *gin.Context) {
	page := model.Page{}
	if err := c.BindJSON(&page); err != nil {
		panic(err)
	}
	prefix := c.Param("prefix")
	service.PagePublishArticleByCategoryPrefix(prefix, &page)
	c.JSON(http.StatusOK, model.Ok(page))
}

func ListRecentlyPublishArticle(c *gin.Context) {
	data := service.ListRecentlyPublishArticle(10)
	c.JSON(http.StatusOK, model.Ok(data))
}
