package router

import (
	"byitter/src/config"
	"byitter/src/middleware/tls"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	r.Use(tls.LoadTls())
	initUserRouters(r)
	if config.C.Debug.Enable {
		r.RunTLS(config.C.App.Addr, "./env/tls/localhost.pem", "./env/tls/localhost-key.pem")
	} else {
		r.Run(config.C.App.Addr)
	}
}
