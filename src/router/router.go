package router

import (
	"byitter/src/config"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	initUserRouters(r)
	r.Run(config.C.App.Addr)
}
