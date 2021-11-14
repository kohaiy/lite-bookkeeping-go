package model

import "github.com/jinzhu/gorm"

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
	Password string
	Slat     string
	Email    string
	Mobile   string
	IsLocked bool
	IsActive bool
}
