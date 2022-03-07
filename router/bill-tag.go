package router

import (
	"github.com/gin-gonic/gin"

	billtag "github.com/kohaiy/lite-bookkeeping-go/module/bill-tag"
)

func UseBillTagRouter(e *gin.Engine) {
	e.GET("/bill-tags", billtag.ListBillTag)
	e.POST("/bill-tag", billtag.AddBillTag)
}
