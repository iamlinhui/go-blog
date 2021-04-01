package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"go-blog/controller"
	"time"
)

func buildRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("web", "web/dist/index.html")
	return p
}

func InitRouter() {
	gin.SetMode("debug")

	engine := gin.New()
	engine.HTMLRender = buildRender()
	engine.Use(gin.Recovery())
	engine.Use(cors.New(
		cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"Content-Length", "text/plain", "Authorization", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		},
	))

	engine.Static("/", "web/dist")
	engine.StaticFile("/favicon.ico", "web/dist/favicon.ico")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(200, "web", nil)
	})

	/*
		前端展示页面接口
	*/
	router := engine.Group("api/v1")
	{
		// 获取个人设置信息
		router.GET("index/:id", controller.Index)
	}

	_ = engine.Run(":8321")

}
