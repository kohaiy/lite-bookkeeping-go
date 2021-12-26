package model

import (
	"github.com/kohaiy/lite-bookkeeping-go/constants"
	"gorm.io/gorm"
)

// id	varchar	32		√
// user_id	varchar	32		√
// bill_account_type_code	int	11	账户类型	√
// name	varchar	50		√
// amount	int	11	金额	√
// remarks	varchar	255	备注

type BillAccount struct {
	gorm.Model
	UserId              uint                      `gorm:"not null"`
	BillAccountTypeCode constants.BillAccountType `gorm:"not null"`
	Name                string                    `gorm:"not null"`
	Amount              int                       `gorm:"not null"`
	Remarks             string
}
