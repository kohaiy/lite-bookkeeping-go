package model

import "gorm.io/gorm"

// id	varchar	32		√
// name	varchar	50		√
// password	varchar	50		√
// slat	varchar	50		√
// email	varchar	50		√
// mobile	varchar	25
// is_lock	bit	1		√
// is_active	bit	1		√

type User struct {
	gorm.Model
	Name     string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
	Slat     string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Mobile   string `gorm:"not null"`
	IsLocked bool   `gorm:"not null;default:false"`
	IsActive bool   `gorm:"not null;default:true"`
}
