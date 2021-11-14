package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/module"
	"github.com/spf13/viper"
)

func NewRouter() *gin.Engine {
	gin.SetMode(viper.GetString("gin.mode"))
	r := gin.Default()

	r.GET("/", module.Index)

	return r
}
