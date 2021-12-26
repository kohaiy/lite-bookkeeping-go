package main

import (
	"fmt"

	"github.com/kohaiy/lite-bookkeeping-go/config"
	"github.com/kohaiy/lite-bookkeeping-go/router"
	"github.com/spf13/viper"

	_ "gorm.io/driver/mysql"
)

func main() {
	config.Init()

	r := router.NewRouter()

	host := viper.GetString("gin.host")
	port := viper.GetString("gin.port")

	if err := r.Run(host + ":" + port); err != nil {
		panic(fmt.Sprintf("Gin 启动失败：%s", err))
	}
}
