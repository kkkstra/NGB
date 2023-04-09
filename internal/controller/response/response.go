package response

import (
	"NGB/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg, omitempty"`
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
	if config.C.Debug.Enable {
		c.JSON(status, Response{
			Code: status,
			Msg:  msg,
			Err:  err,
		})
	}
	c.JSON(status, Response{
		Code: status,
		Msg:  msg,
	})
}
