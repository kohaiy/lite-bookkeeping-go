package config

import (
	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/helper"
	"github.com/spf13/viper"
)

type CommonType struct {
	UniAuthClientId string `json:"uniAuthClientId"`
	UniAuthUrl      string `json:"uniAuthUrl"`
}

func Common(c *gin.Context) {
	res := helper.Res{}

	res.Success(CommonType{
		UniAuthClientId: viper.GetString("uniAuthClientId"),
		UniAuthUrl:      viper.GetString("uniAuthUrl"),
	}).Get(c)
}
