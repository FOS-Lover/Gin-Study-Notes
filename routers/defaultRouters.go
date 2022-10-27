package routers

import (
	"Gin-Note/controllers/index"
	"Gin-Note/middlewares"
	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {

	// 路由分组
	defaultRouters := r.Group("/", middlewares.InitMiddleware)
	// 路由组使用中间间
	defaultRouters.Use(middlewares.InitMiddleware)
	{
		// 自定义控制器抽离
		defaultRouters.GET("/", index.DefaultController{}.Index)
		defaultRouters.GET("/new", index.DefaultController{}.New)
	}
}
