package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/model"
)

type RegisterForm struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Register(c *gin.Context) {

	var form RegisterForm
	if c.ShouldBindJSON(&form) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误。",
		})
		return
	}
	user := &model.User{
		Name: form.Name, Password: form.Password, Slat: "123",
	}
	check := 0
	model.DB.Model(&model.User{}).Where("name=?", form.Name).Count(&check)
	if check > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "用户名已存在。",
		})
		return
	}
	if err := model.DB.Create(user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": user.ID,
	})
}
