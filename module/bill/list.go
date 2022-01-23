package bill

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/constants"
	"github.com/kohaiy/lite-bookkeeping-go/helper"
	"github.com/kohaiy/lite-bookkeeping-go/model"
)

type BillVo struct {
	ID              uint               `json:"id"`
	Amount          int                `json:"amount"`
	ActionTime      time.Time          `json:"actionTime"`
	BillTypeCode    constants.BillType `json:"billTypeCode"`
	BillAccountId   uint               `json:"billAccountId"`
	BillAccountName string             `json:"billAccountName"`
	BillTagId       uint               `json:"billTagId"`
	BillTagName     string             `json:"billTagName"`
	Remarks         string             `json:"remarks"`
	IsIgnore        bool               `json:"isIgnore"`
}

func ListBills(c *gin.Context) {
	res := helper.Res{}
	userId := c.MustGet("UserId").(uint)

	bills := []model.Bill{}
	model.DB.Where("user_id = ?", userId).Order("action_time DESC").Find(&bills)
	formatBills := make([]BillVo, len(bills))
	for index, bill := range bills {
		billAccount := model.BillAccount{}
		model.DB.Where("id = ?", bill.BillAccountId).Find(&billAccount)
		billTag := model.BillTag{}
		model.DB.Where("id = ?", bill.BillTagId).Find(&billTag)
		formatBills[index] = BillVo{
			ID:              bill.ID,
			Amount:          bill.Amount,
			ActionTime:      bill.ActionTime,
			BillTypeCode:    bill.BillTypeCode,
			BillAccountId:   bill.BillAccountId,
			BillAccountName: billAccount.Name,
			BillTagId:       bill.BillTagId,
			BillTagName:     billTag.Name,
			Remarks:         bill.Remarks,
			IsIgnore:        bill.IsIgnore,
		}
	}
	res.Success(formatBills).Get(c)
}
