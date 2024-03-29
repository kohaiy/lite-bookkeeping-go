package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/module/user"
)

func UseUserRouter(e *gin.Engine) {
	r := e.Group("/user")
	r.POST("login", user.Login)
	r.POST("login/oauth", user.LoginOAuth)
	r.POST("bind/oauth", user.BindOAuth)
	r.POST("register", user.Register)
	r.GET("info", user.GetUserInfo)
}
