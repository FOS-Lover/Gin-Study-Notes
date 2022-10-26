# Gin Study Note

### Gin基本使用

- #### 安装
  - `go get github.com/gin-gonic/gin`
  - #### 热加载
    - #### 安装
      - `go install github.com/pilu/fresh@latest`
    - #### 使用
      - `fresh`
- #### 实例
```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 配置路由
	r.GET("/", func(context *gin.Context) {
		context.String(200, "%v", "你好，世界")
	})
	
	// http.StatusOK = 200
	r.GET("/news", func(context *gin.Context) {
		context.String(http.StatusOK, "%v", "主要用于获取数据")
	})
	r.POST("/test", func(context *gin.Context) {
		context.String(200, "%v", "主要用于添加数据")
	})
	r.PUT("/edit", func(context *gin.Context) {
		context.String(200, "%v", "主要用于更新数据")
	})
	r.DELETE("/delete", func(context *gin.Context) {
		context.String(200, "%v", "主要用于删除数据")
	})
	
	// 启动HTTP服务，默认在0.0.0.0:8080 启动服务
	r.Run(":8000") // 参数修改端口
}
```

### 响应数据

- #### String JSON JSONP XML HTML

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	
	// 配置模板文件路径
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "%v", "你好，世界")
	})
	// 1.json map类型响应
	r.GET("/json1", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]interface{}{
			"success": http.StatusOK,
			"data":    "",
		})
	})
	// 2.json gin定义的map类型响应
	r.GET("/json2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"success": http.StatusOK,
			"data":    "",
		})
	})
	// 3.json 结构体类型响应
	type Article struct {
		Title   string `json:"title"`
		Desc    string `json:"desc"`
		Content string `json:"content"`
	}
	r.GET("/json3", func(context *gin.Context) {
		a := &Article{
			Title:   "Test Title",
			Content: "test Content",
			Desc:    "Test Desc",
		}
		context.JSON(http.StatusOK, a)
	})

	// 4.jsonp
	// http://127.0.0.1:8080/jsonp?callback=xxx
	// xxx({"title":"Test Title","desc":"Test Desc","content":"test Content"});
	r.GET("/jsonp", func(context *gin.Context) {
		a := &Article{
			Title:   "Test Title",
			Content: "test Content",
			Desc:    "Test Desc",
		}
		context.JSONP(http.StatusOK, a)
	})

	// 5.xml
	r.GET("/xml", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{
			"success": http.StatusOK,
			"data":    "none",
		})
	})

	// 6.html 渲染模板
	r.GET("/html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "我是后台数据1",
		})
	})

	r.GET("/news", func(context *gin.Context) {
		context.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台数据2",
			"price": 20,
		})
	})

	r.Run(":8080")
}
```

### 模板渲染与模板语法

- #### 书接上文