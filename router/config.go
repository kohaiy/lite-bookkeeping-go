package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/module/config"
)

func UseConfigRouter(e *gin.Engine) {
	e.GET("/config/common", config.Common)
}
