package routers

import (
	"Gin-Note/controllers/admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		// 自定义控制器抽离
		adminRouters.GET("/", admin.IndexController{}.Index)
		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
		adminRouters.GET("/user/edit", admin.UserController{}.Edit)
		adminRouters.GET("/user/del", admin.UserController{}.Delete)
	}
}
