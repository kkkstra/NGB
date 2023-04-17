package main

import (
	"NGB/internal/config"
	"NGB/internal/middleware"
	"NGB/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.JwtAuthMiddleware())
	router.InitRouters(r)
	r.Run(config.C.App.Addr)
}
