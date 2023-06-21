package router

import (
	"NGB/internal/controller"

	"github.com/gin-gonic/gin"
)

func initPostRouters(r *gin.Engine) {
	posts := r.Group("/posts")
	{
		// 搜索帖子
		posts.GET("", controller.GetPostsByKeywords)
		// 新增帖子
		posts.POST("", controller.AddPost)
		postAction := posts.Group("/:post_id")
		{
			// 获取帖子
			postAction.GET("", controller.GetPost)
			// 删除帖子
			postAction.DELETE("", controller.DeletePost)
			// 获取帖子所有点赞者
			postAction.GET("/thumbs", controller.GetAllThumbs)
			// 点赞
			postAction.POST("/thumbs", controller.AddThumbs)
			// 取消点赞
			postAction.DELETE("/thumbs", controller.DeleteThumbs)
		}
	}
}
