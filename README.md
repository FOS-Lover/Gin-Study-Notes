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

- ##### 模板渲染和传值
```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string
	Content string
}

func main() {
	r := gin.Default()

	// 加载模板
	r.LoadHTMLGlob("templates/**/*")

	// 前台
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index/index.html", gin.H{
			"title": "首页",
			"score": 81,
			"hobby": []string{"吃饭", "睡觉", "写代码"},
			"newsList": []interface{}{
				&Article{
					Title:   "新闻标题1",
					Content: "新闻内容1",
				},
				&Article{
					Title:   "新闻标题2",
					Content: "新闻内容2",
				},
			},
			"testSlice": []string{},
			"news": &Article{
				Title:   "新闻标题",
				Content: "新闻内容",
			},
		})
	})
	r.GET("/news", func(context *gin.Context) {
		news := &Article{
			Title:   "新闻标题",
			Content: "新闻内容",
		}
		context.HTML(http.StatusOK, "index/news.html", gin.H{
			"title": "新闻页面",
			"news":  news,
		})
	})

	// 后台
	r.GET("/admin", func(context *gin.Context) {
		context.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})
	r.GET("/admin/news", func(context *gin.Context) {
		context.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "新闻页面",
		})
	})

	r.Run(":8080")
}
```
- ##### html模板语法
```html
{{ define "default/index.html" }}
<!doctype html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport"
        content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>test</title>
</head>
<body>
  <h1>{{.title}}</h1>
<!--  定义变量 -->
  {{ $t := .title }}
  <h4>{{$t}}</h4>
<!-- 条件判断 -->
  {{if ge .score 60}}
  <p>及格</p>
  {{else}}
  <p>不及格</p>
  {{end}}

  {{if gt .score 90}}
  <p>优秀</p>
  {{else if gt .score 80}}
  <p>良好</p>
  {{else if gt .score 60}}
  <p>及格</p>
  {{else}}
  <p>不及格</p>
  {{end}}
<!-- 循环遍历数据 -->
  <ul>
    {{range $key,$value := .hobby}}
    <li>{{$key}}-{{$value}}</li>
    {{end}}
  </ul>

  <br>

  <ul>
    {{range $key,$value := .newsList}}
    <li>{{$key}}-{{$value.Title}}-{{$value.Content}}</li>
    {{end}}
  </ul>

  <br>

  <ul>
    {{range $key,$value := .testSlice}}
    <li>{{$key}}-{{$value.Title}}-{{$value.Content}}</li>
    {{else}}
    <li>切片中没有数据</li>
    {{end}}
  </ul>
<!-- with解构结构体 -->
  <p>{{.news.Title}}</p>
  <p>{{.news.Content}}</p>

  <br>

  {{with .news}}
    <p>{{.Title}}</p>
    <p>{{.Content}}</p>
  {{end}}
</body>
</html>
{{ end }}
```

- #### 书接下文

- ##### 模板渲染和语法

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

type Article struct {
	Title   string
	Content string
}

// 时间戳转换日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func Println(str1 string, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + str2
}

func main() {
	r := gin.Default()

	// 自定义模板函数, 要把函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
		"Println":    Println,
	})

	// 配置静态web目录 第一个参数表示路由，第二个参数表示映射的目录
	r.Static("/static", "./static")

	// 加载模板
	r.LoadHTMLGlob("templates/**/*")

	// 前台
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index/index.html", gin.H{
			"title": "首页",
			"date":  1629423555,
		})
	})
	r.GET("/news", func(context *gin.Context) {
		news := &Article{
			Title:   "新闻标题",
			Content: "新闻内容",
		}
		context.HTML(http.StatusOK, "index/news.html", gin.H{
			"title": "新闻页面",
			"news":  news,
		})
	})
	r.Run(":8080")
}
```

- #### html模板语法
```html
{{ define "public/header.html" }}
  <h1>我是公共头部-{{.title}}</h1>
{{end}}
```

```html
{{ define "default/index.html" }}
<!doctype html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport"
        content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>test</title>
  <link rel="stylesheet" href="/static/css/common.css">
</head>
<body>
<!-- 嵌套使用公共模板 -->
  {{ template "public/header.html" .}}

  <h2>css生效了</h2>
<!--加载图片-->
  <img src="/static/images/yun.jpg" alt="">

<!--  预定义函数 (了解) -->
  <p>{{len .title}}</p>

<!--  自定义模板函数 -->
  <p>{{.date}}</p>
  <p>{{UnixToTime .date}}</p>
  <p>{{Println .title .title}}</p>

</body>
</html>
{{ end }}
```

### 路由

- #### GET POST 传值
```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	// 配置静态web目录 第一个参数表示路由，第二个参数表示映射的目录
	r.Static("/static", "./static")

	// 加载模板
	r.LoadHTMLGlob("templates/**/*")

	// GET请求传值
	r.GET("/", func(context *gin.Context) {
		id := context.Query("id")
		name := context.Query("name")
		page := context.DefaultQuery("page", "1")
		context.JSON(http.StatusOK, gin.H{
			"id":   id,
			"name": name,
			"page": page,
		})
	})
	r.GET("/article", func(context *gin.Context) {
		id := context.DefaultQuery("id", "1") // GET配置默认值
		context.JSON(http.StatusOK, gin.H{
			"id":  id,
			"msg": "新闻详情",
		})
	})

	// POST传值
	r.GET("/user", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index/user.html", gin.H{})
	})
	// 获取post表单数据
	r.POST("/doAddUser", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		age := context.DefaultPostForm("age", "20")	// POST配置默认值
		context.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	r.Run(":8080")
}
```

```html
{{ define "default/user.html" }}
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Title</title>
</head>
<body>
  <form action="/doAddUser" method="post"><br><br>
    用户名: <input type="text" name="username"><br><br>
    密码:   <input type="password" name="password"><br><br>
    <input type="submit" value="提交">
  </form>
</body>
</html>
{{ end }}
```

- #### GET POST 传递的数据绑定到结构体
```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	
	
	// 获取GET POST 传递的数据绑定到结构体
	type UserInfo struct {
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
	}
	r.GET("/getUser", func(context *gin.Context) {
		user := &UserInfo{}
		err := context.ShouldBind(&user)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
		context.JSON(http.StatusOK, user)
	})
	r.POST("/doAddUser", func(context *gin.Context) {
		user := &UserInfo{}
		err := context.ShouldBind(&user)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
		context.JSON(http.StatusOK, user)
	})

	r.Run(":8080")
}
```

- #### 获取 Post Xml数据
```go
package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	// 获取post xml数据
	type Article struct {
		Title   string `json:"title" xml:"title"`
		Content string `json:"content" xml:"content"`
	}
	r.POST("/xml", func(context *gin.Context) {
		var article Article
		data, _ := context.GetRawData() // 读取请求数据
		if err := xml.Unmarshal(data, &article); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		context.JSON(http.StatusOK, article)
	})

	r.Run(":8080")
}
```

- #### 动态路由传值
```go
package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)


func main() {
  r := gin.Default()

  r.GET("/user/:id", func(context *gin.Context) {
    id := context.Param("id")
    context.JSON(http.StatusOK, gin.H{
      "id": id,
    })
  })

  r.Run(":8080")
}
```

- #### 路由分组和路由抽离
```go
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
```

```go
package routers

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func DefaultRoutersInit(r *gin.Engine) {
  // 路由分组
  defaultRouters := r.Group("/")
  {
    defaultRouters.GET("/", func(context *gin.Context) {
      context.HTML(http.StatusOK, "index/index.html", gin.H{})
    })
    defaultRouters.GET("/new", func(context *gin.Context) {
      context.String(200, "new")
    })
  }
}
```

```go
package routers

import "github.com/gin-gonic/gin"

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(context *gin.Context) {
			context.String(200, "api")
		})
		apiRouters.GET("/list", func(context *gin.Context) {
			context.String(200, "list")
		})
	}
}
```

```go
package routers

import "github.com/gin-gonic/gin"

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", func(context *gin.Context) {
			context.String(200, "admin")
		})
		adminRouters.GET("/info", func(context *gin.Context) {
			context.String(200, "info")
		})
	}
}
```

### 控制器

- #### 自定义控制器和控制器分离

- main.go
```go
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
	routers.AdminRoutersInit(r)

	r.Run(":8080")
}
```

- routers
```go
package routers

import (
	"Gin-Note/controllers/index"
	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	// 路由分组
	defaultRouters := r.Group("/")
	{
		// 自定义控制器抽离
		defaultRouters.GET("/", index.DefaultController{}.Index)
		defaultRouters.GET("/new", index.DefaultController{}.New)
	}
}
```
```go
package routers

import (
	"Gin-Note/controllers/admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		// 自定义控制器抽离
		adminRouters.GET("/", admin.IndexController{}.Index)
		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
		adminRouters.GET("/user/edit", admin.UserController{}.Edit)
		adminRouters.GET("/user/del", admin.UserController{}.Delete)
	}
}
```
- controlles
```go
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
```
```go
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
```

### 中间件
  - 中间是匹配路由前和匹配路由完成后执行的操作

```go
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
	

	r.GET("/", initMiddlewareOne, func(context *gin.Context) {
		context.String(200, "123")
	})

	r.GET("/login", initMiddlewareOne, initMiddlewareTwo, func(context *gin.Context) {
		fmt.Println("登录")
		context.String(200, "login")
	})

	r.GET("/logout", func(context *gin.Context) {
		context.String(200, "logout")
	})

	r.Run(":8080")
}
```

- #### 路由里添加中间件

```go
package routers

import (
	"Gin-Note/controllers/index"
	"Gin-Note/middlewares"
	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {

	// 路由分组
	defaultRouters := r.Group("/", middlewares.InitMiddleware)
	// 路由组使用中间间
	defaultRouters.Use(middlewares.InitMiddleware)
	{
		// 自定义控制器抽离
		defaultRouters.GET("/", index.DefaultController{}.Index)
		defaultRouters.GET("/new", index.DefaultController{}.New)
	}
}
```
```go
package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(context *gin.Context) {
	fmt.Println(time.Now())
}
```

- #### 中间件和控制器通信
```go
package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(context *gin.Context) {
	fmt.Println(time.Now())
	context.Set("username", "admin")  // 设置传输数据
}
```

```go
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
```

- #### 中间件分组和协程不影响主程
```go
package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(context *gin.Context) {
	// 定义goroutine统计日志
	cp := context.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("in path: ", cp.Request.URL.Path)
	}()
}
```

```go
package routers

import (
	"Gin-Note/controllers/index"
	"Gin-Note/middlewares"
	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {

	// 路由分组
	defaultRouters := r.Group("/", middlewares.InitMiddleware)
	// 路由组使用中间间
	defaultRouters.Use(middlewares.InitMiddleware)
	{
		// 自定义控制器抽离
		defaultRouters.GET("/", index.DefaultController{}.Index)
		defaultRouters.GET("/new", index.DefaultController{}.New)
	}
}
```

### 自定义Model

```go
package models

import (
	"time"
)

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}
```

### 文件上传

- #### 单文件和多文件

```go
package main

import (
	"Gin-Note/models"
	"Gin-Note/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	
	// 单文件上传
	r.POST("/upload", func(context *gin.Context) {
		file, err := context.FormFile("file")
		// file.Filename 文件名称
		// 文件路径
		dst := path.Join("./static/upload", file.Filename) // ./static/upload/test.jpg
		if err != nil {
			log.Println(err)
		} else {
			// 上传文件
			context.SaveUploadedFile(file, dst)
			context.JSON(http.StatusOK, gin.H{
				"success": true,
				"dst":     dst,
			})
		}
	})

	// 多文件上传
	r.POST("/uploads", func(context *gin.Context) {
		file1, err1 := context.FormFile("file1")
		file2, err2 := context.FormFile("file2")
		// 文件路径
		dst1 := path.Join("./static/upload", file1.Filename)
		if err1 == nil {
			context.SaveUploadedFile(file1, dst1)
		}
		// 文件路径
		dst2 := path.Join("./static/upload", file2.Filename)
		if err2 == nil {
			context.SaveUploadedFile(file2, dst2)
		}
		context.JSON(http.StatusOK, gin.H{
			"success": http.StatusOK,
			"file1":   dst1,
			"file2":   dst2,
		})
	})

	// 相同名多文件上传
	r.POST("/uploadFile", func(context *gin.Context) {
		form, _ := context.MultipartForm()
		files := form.File["files[]"]
		for _, file := range files {
			dst := path.Join("./static/upload", file.Filename)
			context.SaveUploadedFile(file, dst)
		}
		context.JSON(http.StatusOK, gin.H{
			"success": http.StatusOK,
			"data":    [...]string{},
		})
	})

	r.Run(":8080")
}
```

- #### 按日期存储文件

```go
package main

import (
	"Gin-Note/models"
	"Gin-Note/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	

	r.POST("/upload", func(context *gin.Context) {
		// 获取上传文件
		file, err := context.FormFile("file")
		if err == nil {
			// 获取后缀名 判断类型是否正确 .jpg .png .gif .jpeg
			extName := path.Ext(file.Filename)
			allowExtMap := map[string]bool{
				".jpg":  true,
				".png":  true,
				".gif":  true,
				".jpeg": true,
			}
			if _, ok := allowExtMap[extName]; !ok {
				context.String(200, "上传的文件类型不合法")
				return
			}
			// 创建图片保存目录 static/upload/20221028
			day := time.Now().Format("20060102")
			dir := "./static/upload/" + day
			err := os.MkdirAll(dir, 0666)
			if err != nil {
				context.String(200, "创建目录失败")
				return
			}
			// 生成文件名称和文件保存的目录
			fileName := strconv.FormatInt(time.Now().Unix(), 10) + extName
			dst := path.Join(dir, fileName)
			// 执行上传
			context.SaveUploadedFile(file, dst)
			context.JSON(http.StatusOK, gin.H{
				"success": true,
				"dst":     dst,
			})
		}
	})

	r.Run(":8080")
}
```

### Cookie
  - 页面之间数据共享
```go
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

}

func (receiver DefaultController) New(context *gin.Context) {
	// 获取Cookie
	cookie, _ := context.Cookie("username")
	context.String(http.StatusOK, cookie)
}
```

- #### 多个二级域名共享Cookie
  - 域名前面加个`.`
  - ```go
    context.SetCookie("username", "admin", 3600, "/", ".admin.com", false, false)
    
### Session
  - session是另一种记录客户状态的机制，不同的是Cookie保存在客户端浏览器中，而Session保存在服务器

- #### 安装
  - `go get github.com/gin-contrib/sessions`
  - `go get github.com/gin-contrib/sessions/cookie`

- #### 配置和使用

```go
package main

import (
	"Gin-Note/models"
	"Gin-Note/routers"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	

	// 配置session中间件
	// 创建一个基于cookie的存储引擎，secret参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret"))
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))
	
	// 使用session
	r.GET("/session", func(context *gin.Context) {
		// 设置sessions
		session := sessions.Default(context)
		session.Set("username", "test")
		session.Save() // 设置session必须调用
	})
	r.GET("/getSession", func(context *gin.Context) {
		// 获取sessions
		session := sessions.Default(context)
		get := session.Get("username")
		context.String(200, "%v", get)
	})

	r.Run(":8080")
}
```

- #### 分布式Session

- ##### 安装redis
  - `https://github.com/tporadowski/redis/releases`

```go
package main

import (
	"Gin-Note/models"
	"Gin-Note/routers"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

func main() {
	r := gin.Default()
	

	r.Static("/static", "./static")

	// 配置session redis中间件
	store, _ := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("secret"))
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	// 使用session
	r.GET("/session", func(context *gin.Context) {
		// 设置sessions
		session := sessions.Default(context)
		// 配置session过期时间
		session.Options(sessions.Options{
			MaxAge: 3600*6,	// 单位是秒
		})
		session.Set("username", "test")
		session.Save() // 设置session必须调用
	})
	r.GET("/getSession", func(context *gin.Context) {
		// 获取sessions
		session := sessions.Default(context)
		get := session.Get("username")
		context.String(200, "%v", get)
	})

	r.Run(":8080")
}
```

### MySQL数据库

- #### MySQL常用命令
  - `mysql -uroot -p` 连接数据库
  - `show databases;` 查看当前连接的数据库
  - `use gin;` 使用数据库
  - `show tables;` 查看数据库表
  - `select * from users;` 查看表数据
  - `select id,name from users;` 根据字段查找数据
  - `select id,name from users where id=1;` 根据条件查找数据
  - `create database book;` 创建数据库
  - `create table types(id int(11), name varchar(255), number int(3));` 创建表
  - `describe types;` 查看表结构
  - `insert into types(id,name,number) values (1,"Func",1);` 添加数据
  - `update types set number=2 where name="Map";` 修改字段数据
  - `deleter from types where id=2;` 删除数据
  - `select * from types order by id asc;` 以id升序排序
  - `select * from types order by id desc;` 以id降序排序
  - `select * from types order by name desc;` 以name降序排序
  - `select * from types order by name asc;` 以name升序排序
  - `select * from types order by name desc,number asc;` 以name降序和number升序排序
  - `select count(1) from types;` 统计数量
  - `select * from types limit 2;` 查找两条数据
  - `select * from types limit 2,2;` 跳过2条查询2条数据
  - `drop table test;` 删除表
  - `drop database test;` 删除数据库

- #### MySQL关键字和基本操作
  - ##### MySQL字段类型
    - 整数型 `tinyint` `smallint` `mediumint` `int` `bigint`
    - 浮点型 `float` `double` `decimal`
    - 字符型 `char` `varchar`
    - 备注型 `tinytext` `text` `mediumtext` `longtext`
  - ##### 查询语句详解和IN OR AND BETWEEN
    - `select * from class;` 查询所有数据
    - `select name,score from class;` 只查找name,score的数据
    - `select * from class where score > 60;` 查找score大于60的数据
    - `select * from class where email is null;` 查找email为null的数据
    - `select * from class where email is null or email="";` 查找email为null或为""的数据
    - `select * from class where email is not null;` 查找email不为null的数据
    - `select * from class where score >= 60 and score <=90;` 查找score大于等于和60小于等于90的数据
    - `select * from class where score between 60 and 90;`
    - `select * from class where score not between 60 and 90;` 查找score不在大于等于和60小于等于90的数据
    - `select * from class where score=20 or score=30;` 查找score等于20或score等于30的数据
    - `select * from class where score in(20,80,90);` 查找score是20,80,90的数据
    - `select * from class where email like "%test%";` 模糊查找email
  - ##### 分组函数
    - `select avg(score) from class;` 求score平均值
    - `select count(score) from class;` 求score记录总数
    - `select max(score) from class;` 求score最大值
    - `select min(score) from class;` 求score最小值
    - `select sum(score) from class;` 求score总和
    - `select * from class where score in(select max(score) from class);` 查找score最大值且查找对应数据
    - `select * from class where score in(select min(score) from class);` 查找score最小值且查找对应数据
  - ##### 别名
    - `select id,name as n,email as e, score as s from class;`
    - `select min(score) as minscore from class;`

- #### MySQL数据库表关联查询
  - 表与表之间一般存在3种关系，一对一，一对多，多对多关系

- #### MySQL事务和锁定
  - 事务处理可以用来维护数据库的完整性，保证成批的SQL语句要么全部执行，要么全部不执行
    - `begin;` 开启事务
    - `update user set balance=balance-100 where id=1;`
    - `commit;` 提交事务
    - `rollback;` 事务回滚
  - 读锁
    - `lock table user read;` 添加user表为读锁
      - `insert into user(username) values("test2");`
      - `ERROR 1099 (HY000): Table 'user' was locked with a READ lock and can't be updated`
    - `unlock tables;` 释放锁
  - 写锁
    - 只有锁表的用户可以进行读写操作，其他用户不行
    - `lock table user write;` 添加user表为写锁
    - `unlock tables;` 释放锁

### gorm配置和数据库增删改查
  - #### 安装gorm
    - `go get -u gorm.io/gorm`
    - `go get -u gorm.io/driver/mysql`

```go
package admin

import (
	"Gin-Note/datastruct"
	"Gin-Note/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 结构体继承
type UserController struct {
}

func (c UserController) Index(context *gin.Context) {
	// 查询数据库所有数据
	user := []datastruct.User{}
	models.DB.Find(&user)
	context.JSON(http.StatusOK, gin.H{
		"result": user,
	})

	// 查询age大于20的用户
	//user := []datastruct.User{}
	//models.DB.Where("age>20").Find(&user)
	//context.JSON(http.StatusOK, gin.H{
	//	"result": user,
	//})
}

func (c UserController) Add(context *gin.Context) {
	// 添加数据
	user := datastruct.User{
		Username: "test3",
		Age:      27,
		Email:    "test@test.com",
		AddTime:  time.Now().Year(),
	}
	models.DB.Create(&user)
	context.JSON(http.StatusOK, gin.H{
		"result": user,
	})

}

func (c UserController) Edit(context *gin.Context) {
	// 更新数据
	user := datastruct.User{
		Id: 3,
	}
	models.DB.Find(&user).Updates(datastruct.User{
		Username: "ttttttttttt",
	})
	context.JSON(http.StatusOK, gin.H{
		"result": user,
	})
}

func (c UserController) Delete(context *gin.Context) {
	// 删除一条数据
	user := datastruct.User{
		Id: 1,
	}
	models.DB.Find(&user).Delete(&user)
	context.JSON(http.StatusOK, gin.H{
		"result": user,
	})
}
```