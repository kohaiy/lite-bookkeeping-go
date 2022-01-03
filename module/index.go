package module

import (
	"github.com/gin-gonic/gin"
	"github.com/kohaiy/lite-bookkeeping-go/helper"
)

func Index(c *gin.Context) {
	res := helper.Res{}
	res.Message("啥都木有").Get(c)
}
