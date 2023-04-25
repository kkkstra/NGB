package router

import (
	"NGB/internal/controller"

	"github.com/gin-gonic/gin"
)

func initCategoryRouters(r *gin.Engine) {
	categories := r.Group("/categories")
	{
		// 获取所有分类
		categories.GET("", controller.GetAllCategories)
		// 新增分类
		categories.POST("", controller.AddCategory)
		// 获取分类信息
		categories.GET("/:category_id", controller.GetCategory)
		// 删除分类
		categories.DELETE("/:category_id", controller.DeleteCategory)
		// 获取分类下所有帖子，分页处理
		// categories.GET("/:category_id/posts", controller.GetPostsByCategory)
	}
}
