package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/helper"
	"github.com/kohaiy/lite-bookkeeping-go/model"
	"github.com/kohaiy/lite-bookkeeping-go/service"
	"gorm.io/gorm"
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
	var check int64 = 0
	model.DB.Model(&model.User{}).Where("name=?", form.Name).Count(&check)
	if check > 0 {
		res.BadRequest("用户名已存在。").Get(c)
		return
	}
	user.Slat = helper.GenerateSlat()
	user.Password = helper.EncodePassword(form.Password, user.Slat)
	if err := model.DB.Transaction(func(tx *gorm.DB) error {
		err := model.DB.Create(user).Error
		if err == nil {
			err = service.InitBillTag(user.ID, tx).Error
		}
		if err == nil {
			err = service.InitBillAccount(user.ID, tx).Error
		}
		return err
	}); err != nil {
		res.Error(err.Error()).Get(c)
		return
	}

	res.Success(gin.H{
		"id": user.ID,
	}).Get(c)
}
