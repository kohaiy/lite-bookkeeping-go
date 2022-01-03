package router

import (
	"github.com/gin-gonic/gin"

	billtag "github.com/kohaiy/lite-bookkeeping-go/module/bill-tag"
)

func UseBillTagRouter(e *gin.Engine) {
	r := e.Group("/bill-tag")

	r.GET("list", billtag.ListBillTag)
}
