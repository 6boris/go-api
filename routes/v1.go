package routes

import (
	v1 "github.com/dicf/go-api/api/v1"
	"github.com/gin-gonic/gin"
)

func SetupV1(app *gin.Engine) {
	apiv1 := app.Group("/api/v1")

	{
		//	获取单个文章
		apiv1.GET("articles/:id", v1.GetArticle)
		//	获取文章列表
		apiv1.GET("articles", v1.GetArticles)
		//	创建文章
		apiv1.POST("articles", v1.CreateArticle)
		//	更新文章
		apiv1.PUT("articles", v1.UpdateArticle)
		//	删除文章
		apiv1.DELETE("articles/:id", v1.DeleteArticle)

		//	获取单个标签
		apiv1.GET("tags/:id", v1.GetTag)
		//	获取标签
		apiv1.GET("tags", v1.GetTags)
		//	创建标签
		apiv1.POST("tags", v1.CreateTag)
		//	更新标签
		apiv1.PUT("tags/:id", v1.UpdateTag)
		//	删除标签
		apiv1.DELETE("tags/:id", v1.DeleteTag)

	}
}
