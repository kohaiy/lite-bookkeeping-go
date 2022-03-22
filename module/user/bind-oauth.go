package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/constants"
	"github.com/kohaiy/lite-bookkeeping-go/helper"
	"github.com/kohaiy/lite-bookkeeping-go/model"
	"github.com/spf13/viper"
)

type BindOAuthForm struct {
	Type     constants.OauthType `form:"type" binding:"required"`
	Code     string              `form:"code" binding:"required"`
	Name     string              `form:"name" binding:"required"`
	Password string              `form:"password" binding:"required"`
}

func BindOAuth(c *gin.Context) {
	res := &helper.Res{}
	var form BindOAuthForm
	if c.ShouldBindJSON(&form) != nil {
		res.BadRequest("请求参数错误。").Get(c)
		return
	}

	user := &model.User{}

	if rows := model.DB.Where("name=?", form.Name).Find(&user).RowsAffected; rows > 0 {
		password := helper.Md5(helper.Md5(form.Password) + user.Slat)
		if password != user.Password {
			res.BadRequest("用户名或密码错误").Get(c)
			return
		}
	} else {
		res.BadRequest("用户名或密码错误").Get(c)
		return
	}

	keyVal := &model.KeyVal{}
	if rows := model.DB.Where("`key`=?", "user_oauth_"+fmt.Sprint(form.Type)+"_"+form.Code).Where("`expired_at`>=?", helper.GetNowTime()).Find(keyVal).RowsAffected; rows == 0 {
		res.BadRequest("绑定用户失败，请重试").Get(c)
		return
	}

	userInfo := &UserInfo{}
	uniAuthBaseUrl := viper.GetString("uniAuthBaseUrl")
	if err := helper.HttpGetJSON(uniAuthBaseUrl+"/api/user?accessToken="+keyVal.Val, userInfo); err != nil {
		res.BadRequest(err.Error()).Get(c)
		return
	}

	if userInfo.OpenId == "" {
		res.BadRequest("绑定用户失败，请重试").Get(c)
		return
	}

	userOauth := &model.UserOauth{}
	model.DB.Where("code=?", userInfo.OpenId).Where("type=?", form.Type).Find(userOauth)
	userOauth.UserId = user.ID
	userOauth.Type = form.Type
	userOauth.Code = userInfo.OpenId
	if err := model.DB.Save(userOauth).Error; err != nil {
		res.BadRequest(err.Error()).Get(c)
		return
	}
	token := helper.GenerateToken(helper.TokenPayload{
		ID:   user.ID,
		Slat: helper.Md5(user.Slat),
	})
	clientIP := c.MustGet("ClientIP").(string)
	res.Success(gin.H{
		"id":       user.ID,
		"name":     user.Name,
		"email":    user.Email,
		"mobile":   user.Mobile,
		"token":    token,
		"clientIp": clientIP,
	}).Message("Login success").Get(c)
}
