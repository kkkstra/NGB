package main

import (
	"byitter/src/router"
	"byitter/src/util/jwt"
	"github.com/gin-gonic/gin"
)

func main() {
	jwt.InitRSAKey()
	router.InitRouters(gin.Default())
}
