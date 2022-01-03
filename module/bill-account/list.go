package billaccount

import (
	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/helper"
	"github.com/kohaiy/lite-bookkeeping-go/model"
)

type BillAccountVo struct {
	ID                  uint   `json:"id"`
	BillAccountTypeCode uint   `json:"billAccountTypeCode"`
	Name                string `json:"name"`
	Amount              int    `json:"amount"`
	Remarks             string `json:"remarks"`
}

func ListBillAccount(c *gin.Context) {
	res := &helper.Res{}
	userId := c.MustGet("UserId").(uint)
	billAccounts := []model.BillAccount{}
	model.DB.Where("user_id = ?", userId).Order("id").Find(&billAccounts)
	formatBillAccounts := make([]BillAccountVo, len(billAccounts))
	for index, account := range billAccounts {
		formatBillAccounts[index] = BillAccountVo{
			ID:                  account.ID,
			BillAccountTypeCode: uint(account.BillAccountTypeCode),
			Name:                account.Name,
			Amount:              account.Amount,
			Remarks:             account.Remarks,
		}
	}
	res.Success(formatBillAccounts).Get(c)
}
