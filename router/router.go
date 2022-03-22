package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/helper"
	"github.com/kohaiy/lite-bookkeeping-go/middleware"
	"github.com/kohaiy/lite-bookkeeping-go/module"
	"github.com/spf13/viper"
)

type RouteInfo struct {
	Method  string `json:"method"`
	Path    string `json:"path"`
	Handler string `json:"handler"`
}

func NewRouter() *gin.Engine {
	gin.SetMode(viper.GetString("gin.mode"))
	r := gin.Default()

	excludePaths := []string{
		"/",
		"/apis",
		"/user/login",
		"/user/login/oauth",
		"/user/bind/oauth",
		"/user/register",
		"/config/common",
	}
	r.Use(middleware.UseAuth(excludePaths))

	r.GET("/", module.Index)

	UseUserRouter(r)
	UseBillRouter(r)
	UseBillTagRouter(r)
	UseBillAccountRouter(r)
	UseConfigRouter(r)

	rawRoutes := r.Routes()
	routes := make([]RouteInfo, len(rawRoutes))

	for i, r := range rawRoutes {
		routes[i] = RouteInfo{
			Method:  r.Method,
			Path:    r.Path,
			Handler: r.Handler,
		}
	}

	r.GET("/apis", func(c *gin.Context) {
		res := helper.Res{}
		res.Success(routes).Message("显示本系统的所有 API。").Get(c)
	})

	return r
}
