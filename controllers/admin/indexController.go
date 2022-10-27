package admin

import (
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	BaseController
}

func (receiver IndexController) Index(context *gin.Context) {
	//context.String(http.StatusOK, "admin-test")
	// 使用继承
	receiver.Success(context)
}
