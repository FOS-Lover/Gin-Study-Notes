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