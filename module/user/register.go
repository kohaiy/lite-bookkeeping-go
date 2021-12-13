package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/helper"
	"github.com/kohaiy/lite-bookkeeping-go/model"
)

type RegisterForm struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Register(c *gin.Context) {
	res := &helper.Res{}

	var form RegisterForm
	if c.ShouldBindJSON(&form) != nil {
		res.BadRequest("请求参数错误。").Get(c)
		return
	}
	user := &model.User{
		Name: form.Name,
	}
	check := 0
	model.DB.Model(&model.User{}).Where("name=?", form.Name).Count(&check)
	if check > 0 {
		res.BadRequest("用户名已存在。").Get(c)
		return
	}
	user.Slat = helper.GenerateSlat()
	user.Password = helper.EncodePassword(form.Password, user.Slat)
	if err := model.DB.Create(user).Error; err != nil {
		res.Error(err.Error()).Get(c)
		return
	}
	res.Success(gin.H{
		"id": user.ID,
	}).Get(c)
}
