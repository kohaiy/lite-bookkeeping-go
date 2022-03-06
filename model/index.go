package model

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(connectLink string) error {
	db, err := gorm.Open(mysql.Open(connectLink), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("DB connect fail: %s", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	if viper.GetString("gin.mode") == "debug" {
		db.Config.Logger.LogMode(4)
	}
	// db.LogMode(viper.GetString("gin.mode") == "debug")

	DB = db

	autoMigrate()

	return nil
}

func autoMigrate() {
	// 自动迁移表结构
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&UserOauth{})
	DB.AutoMigrate(&BillTag{})
	DB.AutoMigrate(&Bill{})
	DB.AutoMigrate(&BillAccount{})
}
