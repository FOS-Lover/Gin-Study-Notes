package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(context *gin.Context) {
	fmt.Println(time.Now())
	context.Set("username", "admin")

	// 定义goroutine统计日志
	cp := context.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("in path: ", cp.Request.URL.Path)
	}()
}
