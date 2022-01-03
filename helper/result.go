package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Res struct {
	status  int
	code    int
	data    interface{}
	message string
}

func (r *Res) Success(data interface{}) *Res {
	r.Data(data).Code(0).Status(http.StatusOK)
	return r
}

func (r *Res) Error(message string) *Res {
	r.Message(message).Code(-1).Status(http.StatusInternalServerError)
	return r
}

func (r *Res) BadRequest(message string) *Res {
	r.Error(message).Status(http.StatusBadRequest)
	return r
}

func (r *Res) NotFound(message string) *Res {
	r.Error(message).Status(http.StatusNotFound)
	return r
}

func (r *Res) Data(data interface{}) *Res {
	r.data = data
	return r
}

func (r *Res) Code(code int) *Res {
	r.code = code
	return r
}

func (r *Res) Status(status int) *Res {
	r.status = status
	return r
}

func (r *Res) Message(message string) *Res {
	r.message = message
	return r
}

func (r Res) Get(c *gin.Context) {
	c.JSON(r.status, gin.H{
		"code":    r.code,
		"data":    r.data,
		"message": r.message,
	})
}
