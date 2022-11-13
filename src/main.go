package main

import (
	"byitter/src/config"
	"byitter/src/middleware/tls"
	"byitter/src/router"
	"byitter/src/util/initenv"
	"github.com/gin-gonic/gin"
)

func main() {
	initenv.InitEnv()
	r := gin.Default()
	r.Use(tls.LoadTls())
	router.InitRouters(r)
	if config.C.Debug.Enable {
		r.RunTLS(config.C.App.Addr, "./env/tls/localhost.pem", "./env/tls/localhost-key.pem")
	} else {
		r.Run(config.C.App.Addr)
	}
}
