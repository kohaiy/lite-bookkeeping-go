package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/helper"
	"github.com/kohaiy/lite-bookkeeping-go/model"
)

type LoginForm struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	res := &helper.Res{}
	var form LoginForm
	if c.ShouldBindJSON(&form) != nil {
		res.BadRequest("请求参数错误。").Get(c)
		return
	}
	user := &model.User{}
	rows := model.DB.Where("name=?", form.Name).Find(&user).RowsAffected

	if rows > 0 {
		password := helper.Md5(helper.Md5(form.Password) + user.Slat)
		fmt.Println(password + " " + user.Password)
		if password != user.Password {
			res.BadRequest("用户名或密码错误").Get(c)
			return
		}
		token := helper.GenerateToken(helper.TokenPayload{
			ID:   user.ID,
			Slat: helper.Md5(user.Slat),
		})
		res.Success(gin.H{
			"id":     user.ID,
			"name":   user.Name,
			"email":  user.Email,
			"mobile": user.Mobile,
			"token":  token,
		}).Message("Login success").Get(c)
		return
	}
	res.BadRequest("用户名或密码错误").Get(c)
}
