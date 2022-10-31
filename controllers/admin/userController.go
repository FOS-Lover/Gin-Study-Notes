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
