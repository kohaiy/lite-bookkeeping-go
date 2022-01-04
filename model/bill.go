package model

import (
	"time"

	"github.com/kohaiy/lite-bookkeeping-go/constants"
	"gorm.io/gorm"
)

// id	varchar	32		√
// user_id	varchar	32		√
// action_time	datetime		时间	√
// bill_type_code	int	11	类型	√
// bill_account_id	varchar	32	账户	√
// bill_tag_id	varchar	32	标签	√
// amount	int	11	金额	√
// remarks	varchar	255	备注
// is_ignore	bit	1	是否忽略	√

type Bill struct {
	gorm.Model
	UserId        uint               `gorm:"not null"`
	ActionTime    time.Time          `gorm:"not null"`
	BillTypeCode  constants.BillType `gorm:"not null"`
	BillAccountId uint               `gorm:"not null"`
	BillTagId     uint               `gorm:"not null"`
	Amount        int                `gorm:"not null"`
	Remarks       string             `gorm:""`
	IsIgnore      bool               `gorm:"not null;default 0"`
}
