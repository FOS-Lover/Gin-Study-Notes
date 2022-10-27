package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultController struct {
}

func (receiver DefaultController) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "default/index.html", gin.H{})
}

func (receiver DefaultController) New(context *gin.Context) {
	context.String(http.StatusOK, "new")
}
