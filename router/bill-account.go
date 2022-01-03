package router

import (
	"github.com/gin-gonic/gin"

	billaccount "github.com/kohaiy/lite-bookkeeping-go/module/bill-account"
)

func UseBillAccountRouter(e *gin.Engine) {
	r := e.Group("/bill-account")

	r.GET("list", billaccount.ListBillAccount)
}
