package billtag

import (
	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/constants"
	"github.com/kohaiy/lite-bookkeeping-go/helper"
	"github.com/kohaiy/lite-bookkeeping-go/model"
)

type BillTagAddForm struct {
	BillTypeCode constants.BillType `form:"billTypeCode" binding:"required"`
	Name         string             `form:"name" binding:"required"`
	Icon         string             `form:"icon"`
}

func AddBillTag(c *gin.Context) {
	res := helper.Res{}

	userId := c.MustGet("UserId").(uint)

	var form BillTagAddForm
	if !helper.ValidateJSON(&form, c) {
		return
	}
	billTag := &model.BillTag{
		UserId:       userId,
		BillTypeCode: form.BillTypeCode,
		Name:         form.Name,
		Icon:         form.Icon,
	}
	if model.DB.Where("LOWER(name)=LOWER(?)", form.Name).Where("user_id=?", userId).Find(billTag).RowsAffected > 0 {
		res.Error("标签名称已存在").Get(c)
		return
	}

	if err := model.DB.Create(billTag).Error; err != nil {
		res.Error(err.Error()).Get(c)
		return
	}

	res.Success(true).Get(c)
}
