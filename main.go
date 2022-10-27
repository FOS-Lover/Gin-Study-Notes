package main

import (
	"Gin-Note/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")

	r.LoadHTMLGlob("templates/**/*")

	// 路由抽离
	routers.DefaultRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)

	r.Run(":8080")
}
