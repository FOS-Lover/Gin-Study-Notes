package routers

import (
	"Gin-Note/controllers/api"
	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		// 自定义控制器抽离
		apiRouters.GET("/", api.ApiController{}.Index)
		apiRouters.GET("/list", api.ApiController{}.List)
	}
}
