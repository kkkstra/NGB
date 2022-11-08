package router

import (
	"byitter/src/controller"
	"byitter/src/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func initUserRouters(r *gin.Engine) {
	r.POST("/signup", controller.SignUp)
	r.POST("/signin", controller.SignIn)
	user := r.Group("/user")
	user.GET("/:username", controller.GetUserProfile)
	userEdit := user.Group("/:username/edit")
	userEdit.Use(jwt.JwtAuthMiddleware())
	userEdit.POST("/profile", controller.EditUserProfile)
}
