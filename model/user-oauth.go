package model

import "gorm.io/gorm"

type UserOauth struct {
	gorm.Model
	UserId uint   `gorm:"not null"`
	Type   string `gorm:"not null"`
	Code   string `gorm:"not null"`
}
