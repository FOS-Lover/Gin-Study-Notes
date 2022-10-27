package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiController struct {
}

func (receiver ApiController) Index(context *gin.Context) {
	context.String(http.StatusOK, "api")
}

func (receiver ApiController) List(context *gin.Context) {
	context.String(http.StatusOK, "list")
}
