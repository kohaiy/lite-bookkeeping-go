package model

import (
	"github.com/kohaiy/lite-bookkeeping-go/constants"
	"gorm.io/gorm"
)

type UserOauth struct {
	gorm.Model
	UserId uint                `gorm:"not null"`
	Type   constants.OauthType `gorm:"not null"`
	Code   string              `gorm:"not null"`
}
