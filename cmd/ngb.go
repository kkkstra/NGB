package main

import (
	"NGB/internal/config"
	"NGB/internal/middleware"
	"NGB/internal/router"
	"NGB/pkg/logrus"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	logrus.InitLogger(config.C.App.Debug, path.Join(config.C.Log.Filepath, config.C.Log.FilenamePrefix))
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.JwtAuthMiddleware())
	router.InitRouters(r)
	r.Run(config.C.App.Addr)
}
