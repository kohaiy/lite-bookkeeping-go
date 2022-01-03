package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/middleware"
	"github.com/kohaiy/lite-bookkeeping-go/module"
	"github.com/spf13/viper"
)

func NewRouter() *gin.Engine {
	gin.SetMode(viper.GetString("gin.mode"))
	r := gin.Default()

	excludePaths := []string{
		"/user/login",
		"/user/register",
	}
	r.Use(middleware.UseAuth(excludePaths))

	r.GET("/", module.Index)

	UseUserRouter(r)
	UseBillTagRouter(r)
	UseBillAccountRouter(r)

	return r
}
