package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/model"
)

type LoginForm struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var form LoginForm
	if c.ShouldBindJSON(&form) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误。",
		})
		return
	}
	user := &model.User{}
	rows := model.DB.Where("name=?", form.Name).Where("password=?", form.Password).Find(&user).RowsAffected

	if rows > 0 {
		c.JSON(http.StatusOK, gin.H{
			"id":     user.ID,
			"name":   user.Name,
			"email":  user.Email,
			"mobile": user.Mobile,
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "用户名或密码错误",
	})
}
