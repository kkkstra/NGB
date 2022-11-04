package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignUp 用户注册
func SignUp(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}
