package billtag

import (
	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/helper"
	"github.com/kohaiy/lite-bookkeeping-go/model"
)

type BillTagVo struct {
	ID           uint   `json:"id"`
	BillTypeCode uint   `json:"billTypeCode"`
	Name         string `json:"name"`
	Icon         string `json:"icon"`
}

func ListBillTag(c *gin.Context) {
	res := &helper.Res{}
	userId := c.MustGet("UserId").(uint)
	billTags := []model.BillTag{}
	if err := model.DB.Where("user_id = ?", userId).Order("`order` DESC").Order("id").Find(&billTags).Error; err != nil {
		res.Error(err.Error()).Get(c)
		return
	}
	formatBillTags := make([]BillTagVo, len(billTags))
	for index, tag := range billTags {
		formatBillTags[index] = BillTagVo{
			ID:           tag.ID,
			BillTypeCode: uint(tag.BillTypeCode),
			Name:         tag.Name,
			Icon:         tag.Icon,
		}
	}
	res.Success(formatBillTags).Get(c)
}
