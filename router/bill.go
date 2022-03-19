package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/module/bill"
)

func UseBillRouter(e *gin.Engine) {
	e.POST("/bill", bill.AddBill)
	e.GET("/bills", bill.ListBills)
	e.DELETE("/bill/:id", bill.DelBill)
}
