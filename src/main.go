package main

import (
	"byitter/src/controller"
	"byitter/src/model"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	model.ConnectDatabase()
	model.MigrateSchema()

	user := router.Group("/user")
	user.POST("signup", controller.SignUp)

	router.Run(":7777")
}
