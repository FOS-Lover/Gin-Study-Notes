package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 控制器继承
type BaseController struct {
}

func (b BaseController) Success(context *gin.Context) {
	context.String(http.StatusOK, "success")
}

func (b BaseController) Error(context *gin.Context) {
	context.String(http.StatusBadRequest, "error")
}
