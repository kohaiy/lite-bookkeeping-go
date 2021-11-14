package config

import (
	"fmt"
	"os"
	"path"

	"github.com/kohaiy/lite-bookkeeping-go/model"
	"github.com/spf13/viper"
)

func Init() {
	dir, _ := os.Getwd()
	configPath := path.Join(dir, "config")
	configName := "config"
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("配置文件读取失败：%s", err))
	}

	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	database := viper.GetString("mysql.database")

	connectLink := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	if err := model.ConnectDB(connectLink); err != nil {
		panic(err)
	}
}
