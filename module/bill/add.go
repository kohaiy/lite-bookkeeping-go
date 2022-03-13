package bill

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/helper"
	"github.com/kohaiy/lite-bookkeeping-go/model"
	"gorm.io/gorm"
)

type BillAddForm struct {
	ActionTime    time.Time `form:"actionTime" binding:"required"`
	BillAccountId uint      `form:"billAccountId" binding:"required"`
	BillTagId     uint      `form:"billTagId" binding:"required"`
	Amount        int       `form:"amount" binding:"required"`
	Remarks       string    `form:"remarks" binding:""`
	IsIgnore      bool      `form:"isIgnore" binding:""`
}

func AddBill(c *gin.Context) {
	res := helper.Res{}
	userId := c.MustGet("UserId").(uint)

	var form BillAddForm
	if !helper.ValidateJSON(&form, c) {
		return
	}

	billAccount := model.BillAccount{}
	if model.DB.Where("user_id = ?", userId).Where("id = ?", form.BillAccountId).Find(&billAccount).RowsAffected <= 0 {
		res.BadRequest("所选账户不存在。").Get(c)
		return
	}

	billTag := model.BillTag{}
	if model.DB.Where("user_id = ?", userId).Where("id = ?", form.BillTagId).Find(&billTag).RowsAffected <= 0 {
		res.BadRequest("所选标签不存在。").Get(c)
		return
	}

	bill := &model.Bill{
		UserId:        userId,
		ActionTime:    form.ActionTime,
		BillTypeCode:  billTag.BillTypeCode,
		BillAccountId: form.BillAccountId,
		BillTagId:     form.BillTagId,
		Amount:        form.Amount,
		Remarks:       form.Remarks,
		IsIgnore:      form.IsIgnore,
	}
	if err := model.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(bill).Error; err != nil {
			return err
		}
		billTag.Order += 1
		if err := tx.Save(billTag).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		res.Error(err.Error()).Get(c)
		return
	}

	res.Success(bill.ID).Get(c)
}
