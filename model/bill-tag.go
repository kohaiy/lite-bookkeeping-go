package model

import (
	"github.com/kohaiy/lite-bookkeeping-go/constants"
	"gorm.io/gorm"
)

// id	varchar	32		√
// user_id	varchar	32		√
// bill_type_code	int	11	类型	√
// name	varchar	50		√
// icon	varchar	50

type BillTag struct {
	gorm.Model
	UserId       uint               `gorm:"not null"`
	BillTypeCode constants.BillType `gorm:"not null"`
	Name         string             `gorm:"not null"`
	Icon         string             `gorm:""`
	Order        uint               `gorm:""`
}
