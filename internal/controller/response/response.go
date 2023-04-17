package response

import (
	"NGB/internal/config"
	"NGB/pkg/logrus"
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
	if config.C.App.Debug {
		c.JSON(status, Response{
			Code: status,
			Msg:  msg,
			Err:  err,
		})
	} else {
		c.JSON(status, Response{
			Code: status,
			Msg:  msg,
		})
	}
	logrus.Logger.Error(err)
}
