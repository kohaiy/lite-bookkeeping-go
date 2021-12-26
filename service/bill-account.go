package service

import (
	"github.com/kohaiy/lite-bookkeeping-go/constants"
	"github.com/kohaiy/lite-bookkeeping-go/model"
	"gorm.io/gorm"
)

func InitBillAccount(userId uint, tx *gorm.DB) *gorm.DB {
	billAccounts := []model.BillAccount{
		{Name: "现金", BillAccountTypeCode: constants.BAT_CASH, UserId: userId},
		{Name: "支付宝", BillAccountTypeCode: constants.BAT_IDEAL, UserId: userId},
		{Name: "微信钱包", BillAccountTypeCode: constants.BAT_IDEAL, UserId: userId},
		{Name: "工商银行卡", BillAccountTypeCode: constants.BAT_CARD, UserId: userId},
	}
	return tx.Create(&billAccounts)
}
