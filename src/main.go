package main

import (
	"byitter/src/router"
	"byitter/src/util/Init"
	"github.com/gin-gonic/gin"
)

func main() {
	Init.InitEnv()
	router.InitRouters(gin.Default())
}
