package main

import (
	"byitter/src/router"
	"github.com/gin-gonic/gin"
)

func main() {
	router.InitRouters(gin.Default())
}
