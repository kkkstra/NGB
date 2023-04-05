package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
	Err  []string    `json:"error,omitempty"`
}

func Success(c *gin.Context, status int, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: status,
		Data: data,
		Msg:  msg,
	})
}

func Error(c *gin.Context, status int, msg string, err ...string) {
	c.JSON(status, Response{
		Code: status,
		Msg:  msg,
		Err:  err,
	})
}
