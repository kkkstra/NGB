package main

import (
	"NGB/internal/config"
	"NGB/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouters(r)
	r.Run(config.C.App.Addr)
}
