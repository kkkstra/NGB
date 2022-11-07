package main

import (
	"byitter/src/jwt"
	"byitter/src/router"
	"github.com/gin-gonic/gin"
)

func main() {
	jwt.InitRSAKey()
	router.InitRouters(gin.Default())
}
