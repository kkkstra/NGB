package router

import (
	"NGB/internal/controller"
	"NGB/internal/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func initUserRouters(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/signup", controller.SignUp)
		user.POST("/signin", controller.SignIn)
		user.GET("/:username", controller.GetUserProfile)
		userAction := user.Group("/:username")
		{
			userAction.Use(jwt.JwtAuthMiddleware())
			userActionEdit := userAction.Group("/edit")
			{
				userActionEdit.PUT("/profile", controller.EditUserProfile)
				userActionEdit.PUT("/password", controller.EditUserPassword)
				userActionEdit.PUT("/email", controller.EditUserEmail)
			}
			userAction.DELETE("/delete", controller.DeleteUser)
		}
	}
}
