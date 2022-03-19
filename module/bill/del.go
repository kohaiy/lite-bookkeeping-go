package bill

import (
	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/helper"
	"github.com/kohaiy/lite-bookkeeping-go/model"
)

type BillDeleteForm struct {
	Id uint `uri:"id" binding:"required"`
}

func DelBill(c *gin.Context) {
	res := helper.Res{}
	userId := c.MustGet("UserId").(uint)

	var form BillDeleteForm
	if !helper.ValidateUri(&form, c) {
		return
	}

	bill := model.Bill{}
	if model.DB.Where("user_id = ?", userId).Where("id = ?", form.Id).Find(&bill).RowsAffected <= 0 {
		res.BadRequest("账单不存在。").Get(c)
		return
	}
	model.DB.Delete(&bill)
	res.Success(bill.ID).Get(c)
}
