package main

import (
	"Gin-Note/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 中间件
func initMiddlewareOne(context *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("我先调用-initMiddlewareOne")

	// 调用该请求的剩余处理程序
	context.Next()

	end := time.Now().UnixNano()
	fmt.Println("我最后调用-initMiddlewareOne")

	fmt.Println("TimeOne: ", end-start)
}

func initMiddlewareTwo(context *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("我先调用-initMiddlewareTwo")

	// 调用该请求的剩余处理程序
	context.Next()

	end := time.Now().UnixNano()
	fmt.Println("我最后调用-initMiddlewareTwo")

	fmt.Println("TimeTwo: ", end-start)
}

func main() {
	r := gin.Default()

	r.Static("/static", "./static")

	r.LoadHTMLGlob("templates/**/*")

	// 配置全局中间件
	r.Use(initMiddlewareOne, initMiddlewareOne)

	// 路由抽离
	routers.DefaultRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)

	r.Run(":8080")
}
