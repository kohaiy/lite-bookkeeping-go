package helper

import "github.com/gin-gonic/gin"

func ValidateJSON(obj interface{}, c *gin.Context) bool {
	res := Res{}
	if err := c.ShouldBindJSON(obj); err != nil {
		e := err.Error()
		res.BadRequest("请求参数错误。").Data(e).Get(c)
		return false
	}
	return true
}

func ValidateQuery(obj interface{}, c *gin.Context) bool {
	res := Res{}
	if err := c.ShouldBindQuery(obj); err != nil {
		e := err.Error()
		res.BadRequest("请求参数错误。").Data(e).Get(c)
		return false
	}
	return true
}

func ValidateUri(obj interface{}, c *gin.Context) bool {
	res := Res{}
	if err := c.ShouldBindUri(obj); err != nil {
		e := err.Error()
		res.BadRequest("请求参数错误。").Data(e).Get(c)
		return false
	}
	return true
}
