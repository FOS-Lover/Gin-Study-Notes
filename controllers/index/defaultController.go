package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultController struct {
}

func (receiver DefaultController) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "default/index.html", gin.H{})
	value, exists := context.Get("username") // 获取中间件设置的数据
	v, _ := value.(string)                   // 类型断言
	fmt.Println(value, exists, v)
}

func (receiver DefaultController) New(context *gin.Context) {
	context.String(http.StatusOK, "new")
}
