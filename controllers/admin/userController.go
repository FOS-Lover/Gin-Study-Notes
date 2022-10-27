package admin

import "github.com/gin-gonic/gin"

// 结构体继承
type UserController struct {
}

func (c UserController) Index(context *gin.Context) {
	context.String(200, "user")
}

func (c UserController) Add(context *gin.Context) {
	context.String(200, "add")
}

func (c UserController) Edit(context *gin.Context) {
	context.String(200, "edit")
}

func (c UserController) Delete(context *gin.Context) {
	context.String(200, "del")
}
