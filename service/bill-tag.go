package service

import (
	"github.com/kohaiy/lite-bookkeeping-go/constants"
	"github.com/kohaiy/lite-bookkeeping-go/model"
	"gorm.io/gorm"
)

func InitBillTag(userId uint, tx *gorm.DB) *gorm.DB {
	billTags := []model.BillTag{
		{Name: "薪资", BillTypeCode: constants.BT_INCOME, UserId: userId},
		{Name: "收益", BillTypeCode: constants.BT_INCOME, UserId: userId},
		{Name: "餐饮", BillTypeCode: constants.BT_PAY, UserId: userId},
		{Name: "购物", BillTypeCode: constants.BT_PAY, UserId: userId},
		{Name: "房租", BillTypeCode: constants.BT_PAY, UserId: userId},
	}
	return tx.Create(&billTags)
}
