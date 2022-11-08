package router

import (
	"byitter/src/controller"
	"github.com/gin-gonic/gin"
)

func initUserRouters(r *gin.Engine) {
	user := r.Group("/user")
	user.POST("signup", controller.SignUp)
	user.GET("signin", controller.SignIn)
}
