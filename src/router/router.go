package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	initUserRouters(r)
}
