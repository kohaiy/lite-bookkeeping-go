package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/helper"
	"github.com/kohaiy/lite-bookkeeping-go/model"
)

func GetUserInfo(c *gin.Context) {
	res := helper.Res{}
	userId := c.MustGet("UserId").(uint)

	user := &model.User{}
	rows := model.DB.Where("id=?", userId).Find(&user).RowsAffected

	if rows > 0 {
		res.Success(gin.H{
			"id":     user.ID,
			"name":   user.Name,
			"email":  user.Email,
			"mobile": user.Mobile,
		}).Message("Login success").Get(c)
		return
	}
	res.Unauthorized("").Get(c)
}
