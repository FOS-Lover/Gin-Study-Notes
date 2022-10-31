package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultController struct {
}

func (receiver DefaultController) Index(context *gin.Context) {
	// 设置Cookie 参数: 属性 值 过期时间 域 域名 设置http/https 反正xss攻击
	context.SetCookie("username", "admin", 3600, "/", "localhost", false, false)
	// 删除Cookie
	//context.SetCookie("username", "admin", -1, "/", "localhost", false, false)

	context.HTML(http.StatusOK, "default/index.html", gin.H{})
	value, exists := context.Get("username") // 获取中间件设置的数据
	v, _ := value.(string)                   // 类型断言
	fmt.Println(value, exists, v)

}

func (receiver DefaultController) New(context *gin.Context) {
	// 获取Cookie
	cookie, _ := context.Cookie("username")
	context.String(http.StatusOK, cookie)
}
