package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags 文章管理
// @Summary 获取单篇文章
// @Description 根据传入的ID获取文章
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Success 200
// @Router /api/v1/articles/{id} [get]
func GetArticle(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": nil,
	})
}

// @Tags 文章管理
// @Summary 获取文章列表
// @Description 根据传入的参数筛选文章
// @Accept  json
// @Produce  json
// @Param title query int false "Title"
// @Success 200
// @Router /api/v1/articles [get]
func GetArticles(c *gin.Context) {
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": title,
	})
}

// @Tags 文章管理
// @Summary 创建文章
// @Description 根据传入的参数创建文章
// @Accept  json
// @Produce  json
// @Param title query int false "Title"
// @Success 200
// @Router /api/v1/articles [post]
func CreateArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "CreateArticle",
		"data": nil,
	})
}

// @Tags 文章管理
// @Summary 修改文章
// @Description 根据传入的参数修改文章
// @Accept  json
// @Produce  json
// @Param title query int false "Title"
// @Success 200
// @Router /api/v1/articles [put]
func UpdateArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "UpdateArticle",
		"data": nil,
	})
}

// @Tags 文章管理
// @Summary 删除文章
// @Description 根据传入的参数删除文章
// @Accept  json
// @Produce  json
// @Param title query int false "Title"
// @Success 200
// @Router /api/v1/articles [delete]
func DeleteArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "DeleteArticle",
		"data": nil,
	})
}
