package routes

import (
	"blog/server/config"
	"blog/server/controller"
	"blog/server/model"
	"blog/server/util"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
	"net/http"
	"reflect"
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
	RegisterValidate()

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
	engine.Static("/static", "web/dist/static")
	engine.StaticFile("/favicon.ico", "web/dist/favicon.ico")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(200, "web", nil)
	})
	/*
		前端展示页面接口
	*/
	publicRouter := engine.Group("/api/v1")
	{
		// 首页分页展示
		publicRouter.POST("/page", controller.PagePublishArticle)
		// 所有文章分类标签
		publicRouter.POST("/category", controller.ListCategory)
		// 最近文章列表
		publicRouter.POST("/recently", controller.ListRecentlyPublishArticle)

		// 文章详情
		publicRouter.POST("/article/:url", controller.GetPublishArticleByUrl)
		// 根据分类分页展示
		publicRouter.POST("/page/category/:prefix", controller.PagePublishArticleByCategoryPrefix)
	}

	privateRouter := engine.Group("/api/v1/admin")
	privateRouter.Use(util.JwtAuth())
	{

	}

	engine.NoRoute(func(context *gin.Context) {
		// context.Redirect(302,"/") 重定向到主页
		context.JSON(http.StatusOK, model.GetException(model.UrlNotFound))
	})
	return engine
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
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
				c.JSON(http.StatusOK, model.BuildException(model.Error, err.Error()))
			}
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

func RegisterValidate() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册 model.LocalTime 类型的自定义校验规则
		v.RegisterCustomTypeFunc(ValidateJSONDateType, model.LocalTime{})
	}
}

func ValidateJSONDateType(field reflect.Value) interface{} {
	if field.Type() == reflect.TypeOf(model.LocalTime{}) {
		timeStr := field.Interface().(model.LocalTime).String()
		// 0001-01-01 00:00:00 是 go 中 time.Time 类型的空值
		// 这里返回 Nil 则会被 validator 判定为空值，而无法通过 `binding:"required"` 规则
		if timeStr == "0001-01-01 00:00:00" {
			return nil
		}
		return timeStr
	}
	return nil
}
