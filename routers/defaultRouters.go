package routers

import (
	"Gin-Note/controllers/index"
	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	// 路由分组
	defaultRouters := r.Group("/")
	{
		// 自定义控制器抽离
		defaultRouters.GET("/", index.DefaultController{}.Index)
		defaultRouters.GET("/new", index.DefaultController{}.New)
	}
}
