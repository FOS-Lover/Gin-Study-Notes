package main

import (
	"Gin-Note/models"
	"Gin-Note/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
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

	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})

	r.Static("/static", "./static")

	r.LoadHTMLGlob("templates/**/*")

	// 配置全局中间件
	r.Use(initMiddlewareOne, initMiddlewareOne)
	//
	//// 配置session中间件
	//store, _ := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("secret"))
	//// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	//r.Use(sessions.Sessions("mysession", store))
	//
	//// 使用session
	//r.GET("/session", func(context *gin.Context) {
	//	// 设置sessions
	//	session := sessions.Default(context)
	//	// 配置session过期时间
	//	session.Options(sessions.Options{
	//		MaxAge: 3600 * 6, // 单位是秒
	//	})
	//	session.Set("username", "test")
	//	session.Save() // 设置session必须调用
	//})
	//r.GET("/getSession", func(context *gin.Context) {
	//	// 获取sessions
	//	session := sessions.Default(context)
	//	get := session.Get("username")
	//	context.String(200, "%v", get)
	//})

	// 路由抽离
	routers.DefaultRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)

	r.Run(":8080")
}
