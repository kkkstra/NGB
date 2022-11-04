package main

import (
	"byitter/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.GET("register", controller.SignUp)
	}

	_ = router.Run(":7777")
}
