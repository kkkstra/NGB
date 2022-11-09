package main

import (
	"byitter/src/router"
	"byitter/src/util/initEnv"
	"github.com/gin-gonic/gin"
)

func main() {
	initEnv.InitEnv()
	router.InitRouters(gin.Default())
}
