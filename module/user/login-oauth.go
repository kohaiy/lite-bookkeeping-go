package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/helper"
	"github.com/spf13/viper"
)

type LoginOAuthForm struct {
	Type string `form:"type" binding:"required"`
	Code string `form:"code" binding:"required"`
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
	clientSecret := viper.GetString("uniAuthSecret")
	uniAuthClientId := viper.GetString("uniAuthClientId")
	data := make(map[string]interface{})
	data["clientId"] = uniAuthClientId
	data["clientSecret"] = clientSecret
	data["code"] = form.Code
	bytesData, _ := json.Marshal(data)
	resp, _ := http.Post("https://testing-auth.kohai.dev/api/oauth/access-token", "application/json", bytes.NewReader(bytesData))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	accessTokenJSON := &AccessTokenJSON{}
	json.Unmarshal(body, accessTokenJSON)
	fmt.Println(accessTokenJSON)

	if accessTokenJSON.AccessToken == "" {
		res.BadRequest("授权登录失败，请重试").Get(c)
		return
	}

	res2, _ := http.Get("https://testing-auth.kohai.dev/api/user?accessToken=" + accessTokenJSON.AccessToken)
	defer res2.Body.Close()
	body2, _ := ioutil.ReadAll(res2.Body)

	userInfo := &UserInfo{}
	json.Unmarshal(body2, userInfo)

	if userInfo.OpenId == "" {
		res.BadRequest("授权登录失败，请重试").Get(c)
		return
	}

	fmt.Println(userInfo)
	res.Success(userInfo).Get(c)
	// userOauth := &model.UserOauth{}
	// rows := model.DB.Where("code")
	// user := &model.User{}
	// rows := model.DB.Where("name=?", form.Name).Find(&user).RowsAffected

	// if rows > 0 {
	// 	password := helper.Md5(helper.Md5(form.Password) + user.Slat)
	// 	if password != user.Password {
	// 		res.BadRequest("用户名或密码错误").Get(c)
	// 		return
	// 	}
	// 	token := helper.GenerateToken(helper.TokenPayload{
	// 		ID:   user.ID,
	// 		Slat: helper.Md5(user.Slat),
	// 	})
	// 	clientIP := c.MustGet("ClientIP").(string)
	// 	res.Success(gin.H{
	// 		"id":       user.ID,
	// 		"name":     user.Name,
	// 		"email":    user.Email,
	// 		"mobile":   user.Mobile,
	// 		"token":    token,
	// 		"clientIp": clientIP,
	// 	}).Message("Login success").Get(c)
	// 	return
	// }
	// res.BadRequest("用户名或密码错误").Get(c)
}
