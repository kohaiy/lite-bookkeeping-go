package model

import (
	"time"

	"gorm.io/gorm"
)

type KeyVal struct {
	gorm.Model
	Key       string    `gorm:"not null"`
	Val       string    `gorm:""`
	ExpiredAt time.Time `gorm:""`
}
