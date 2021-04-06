package routes

import (
	"blog/src/config"
	"blog/src/controller"
	"blog/src/model"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

func buildRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("web", "web/dist/index.html")
	return p
}

func InitRouter() *gin.Engine {
	gin.SetMode(config.AppMode)

	engine := gin.New()
	engine.HTMLRender = buildRender()
	engine.Use(Recover)
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
	engine.Static("/css", "web/dist/css")
	engine.Static("/img", "web/dist/img")
	engine.Static("/js", "web/dist/js")
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

	engine.NoRoute(func(context *gin.Context) {
		// context.Redirect(302,"/") 重定向到主页
		context.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "请求路径不正确",
		})
	})

	return engine
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			if _, ok := r.(model.Exception); ok {
				c.JSON(http.StatusOK, r)
				return
			}
			if _, ok := r.(*model.Exception); ok {
				c.JSON(http.StatusOK, r)
				return
			}
			if err, ok := r.(error); ok {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					c.JSON(http.StatusOK, model.GetException(model.NotFound))
					return
				}
				//封装通用json返回
				c.JSON(http.StatusOK, gin.H{
					"code": model.Error,
					"msg":  err.Error(),
				})
			}
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}
