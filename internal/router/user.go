package router

import (
	"NGB/internal/controller"

	"github.com/gin-gonic/gin"
)

func initUserRouters(r *gin.Engine) {
	users := r.Group("/users")
	{
		// 注册
		users.POST("", controller.SignUp)
		userAction := users.Group("/:username")
		{
			// 获取用户资料
			userAction.GET("/profile", controller.GetUserProfile)
			// 更新用户资料
			userAction.POST("/profile", controller.EditUserProfile)
			// 更新密码
			userAction.POST("/password", controller.EditUserPassword)
			// 更新邮箱
			userAction.POST("/email", controller.EditUserEmail)
			// 删除用户
			userAction.DELETE("", controller.DeleteUser)
			// 获取关注
			userAction.GET("/following", controller.GetAllFollowings)
			// 新增关注
			userAction.POST("/following", controller.AddFollowing)
			// 取消关注
			userAction.DELETE("/following", controller.DeleteFollowing)
		}
	}
	session := r.Group("/session")
	{
		// 登录
		session.POST("", controller.SignIn)
		session.GET("/email", controller.GetSignInCode)
	}
}
