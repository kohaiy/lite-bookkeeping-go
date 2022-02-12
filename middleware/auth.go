package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/helper"
	"github.com/kohaiy/lite-bookkeeping-go/model"
)

func UseAuth(excludePaths []string) gin.HandlerFunc {
	return func(c *gin.Context) {

		clientIP := c.GetHeader("HTTP_CF_CONNECTING_IP")
		if clientIP == "" {
			clientIP = c.GetHeader("REMOTE_ADDR")
		}
		if clientIP == "" {
			clientIP = c.ClientIP()
		}
		fmt.Println("HTTP_CF_CONNECTING_IP: " + c.GetHeader("HTTP_CF_CONNECTING_IP") + " / " + c.ClientIP())
		fmt.Println("REMOTE_ADDR: " + c.GetHeader("REMOTE_ADDR") + " / " + c.ClientIP())
		c.Set("ClientIP", clientIP)

		for _, p := range excludePaths {
			if c.Request.URL.Path == p {
				c.Next()
				return
			}
		}
		tokenString := c.GetHeader("authorization")
		token := helper.ParseToken(tokenString)
		res := &helper.Res{}
		if token == nil {
			res.Unauthorized("用户验证出错了").Get(c)
			c.Abort()
			return
		}
		user := &model.User{}
		if rows := model.DB.Where("id=?", token.ID).Find(&user).RowsAffected; rows <= 0 {
			res.NotFound("用户不存在").Get(c)
			c.Abort()
			return
		}
		if helper.Md5(user.Slat) != token.Slat {
			res.Unauthorized("用户验证出错了").Get(c)
			c.Abort()
			return
		}
		c.Set("UserId", token.ID)

		c.Next()
	}
}
