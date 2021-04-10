package controller

import (
	"blog/server/model"
	"blog/server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCategory(c *gin.Context) {
	data := service.ListCategory()
	c.JSON(http.StatusOK, model.Ok(data))
}
