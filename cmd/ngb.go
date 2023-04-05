package main

import (
	"NGB/internal/config"
	"NGB/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.Use(tls.LoadTls())
	router.InitRouters(r)
	// if config.C.Debug.Enable {
	// 	r.RunTLS(config.C.App.Addr, "./env/tls/localhost.pem", "./env/tls/localhost-key.pem")
	// } else {
	r.Run(config.C.App.Addr)
	// }
}
