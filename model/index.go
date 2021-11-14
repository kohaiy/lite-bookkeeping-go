package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func ConnectDB(connectLink string) error {
	db, err := gorm.Open("mysql", connectLink)
	if err != nil {
		return fmt.Errorf("DB connect fail: %s", err)
	}
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)

	db.LogMode(viper.GetString("gin.mode") == "debug")

	DB = db

	autoMigrate()

	return nil
}

func autoMigrate() {
	// 自动迁移表结构
	DB.AutoMigrate(&User{})
}
