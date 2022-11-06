package router

import (
	"byitter/src/config"
	"byitter/src/middleware/tls"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	r.Use(tls.LoadTls())
	initUserRouters(r)
	r.RunTLS(config.C.App.Addr, "./env/tls/localhost.pem", "./env/tls/localhost-key.pem")
}
