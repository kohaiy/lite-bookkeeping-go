package router

import (
	"github.com/gin-gonic/gin"

	billaccount "github.com/kohaiy/lite-bookkeeping-go/module/bill-account"
)

func UseBillAccountRouter(e *gin.Engine) {
	e.GET("/bill-accounts", billaccount.ListBillAccount)
}
