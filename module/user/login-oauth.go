package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/constants"
	"github.com/kohaiy/lite-bookkeeping-go/helper"
	"github.com/kohaiy/lite-bookkeeping-go/model"
	"github.com/spf13/viper"
)

type LoginOAuthForm struct {
	Type constants.OauthType `form:"type" binding:"required"`
	Code string              `form:"code" binding:"required"`
}

type AccessTokenJSON struct {
	AccessToken string `json:"accessToken"`
}

type UserInfo struct {
	OpenId   string `json:"openId"`
	UserName string `json:"userName"`
}

func LoginOAuth(c *gin.Context) {
	res := &helper.Res{}
	var form LoginOAuthForm
	if c.ShouldBindJSON(&form) != nil {
		res.BadRequest("请求参数错误。").Get(c)
		return
	}

	accessToken, err := helper.GetUniAuthAccessToken(form.Code)
	if err != nil {
		res.BadRequest(err.Error()).Get(c)
		return
	}

	userInfo := &UserInfo{}
	uniAuthBaseUrl := viper.GetString("uniAuthBaseUrl")
	if err := helper.HttpGetJSON(uniAuthBaseUrl+"/api/user?accessToken="+accessToken, userInfo); err != nil {
		res.BadRequest(err.Error()).Get(c)
		return
	}

	if userInfo.OpenId == "" {
		res.BadRequest("授权登录失败，请重试").Get(c)
		return
	}

	fmt.Println(userInfo)
	userOauth := &model.UserOauth{}
	if rows := model.DB.Where("code=?", userInfo.OpenId).Find(userOauth).RowsAffected; rows == 0 {
		keyVal := &model.KeyVal{
			Key:       "user_oauth_" + fmt.Sprint(form.Type) + "_" + form.Code,
			Val:       accessToken,
			ExpiredAt: helper.GetNowTime().Add(constants.UserOauthExpiredAt),
		}
		if err := model.DB.Save(keyVal).Error; err != nil {
			res.BadRequest(err.Error()).Get(c)
			return
		}
		res.Success(gin.H{
			"isBind": false,
		}).Get(c)
		return
	}

	user := &model.User{}
	rows := model.DB.Where("id=?", userOauth.UserId).Find(&user).RowsAffected

	if rows > 0 {
		token := helper.GenerateToken(helper.TokenPayload{
			ID:   user.ID,
			Slat: helper.Md5(user.Slat),
		})
		clientIP := c.MustGet("ClientIP").(string)
		res.Success(gin.H{
			"isBind": true,
			"data": gin.H{
				"id":       user.ID,
				"name":     user.Name,
				"email":    user.Email,
				"mobile":   user.Mobile,
				"token":    token,
				"clientIp": clientIP,
			},
		}).Message("Login success").Get(c)
		return
	}
	res.BadRequest("用户名不存在").Get(c)
}
