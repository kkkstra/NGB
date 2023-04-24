package router

import (
	"NGB/internal/controller"

	"github.com/gin-gonic/gin"
)

func initPostRouters(r *gin.Engine) {
	posts := r.Group("/posts")
	{
		postAction := posts.Group("/:post_id")
		{
			// 获取帖子
			postAction.GET("", controller.GetPost)
			postAction.POST("", controller.AddPost)
			postAction.DELETE("", controller.DeletePost)
		}
	}
}
