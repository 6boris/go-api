package v1

import (
	"github.com/dicf/go-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Tags 标签管理
// @Summary 获取单个标签
// @Description 根据传入的ID获取标签
// @Accept  json
// @Produce  json
// @Param id path int true "标签ID"
// @Success 200
// @Router /api/v1/tags/{id} [get]
func GetTag(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	tag := service.GetTag(id)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "GetTag",
		"data": tag,
	})
}

// @Tags 标签管理
// @Summary 获取标签列表
// @Description 根据传入的参数筛选标签
// @Accept  json
// @Produce  json
// @Param name query string false "标签名称"
// @Success 200
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {

	filter := make(map[string]interface{})
	filter["name"] = c.Query("name")

	tags := service.GetTags(filter)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "GetTags",
		"data": tags,
	})
}

// @Tags 标签管理
// @Summary 创建标签
// @Description 根据传入的参数创建标签
// @Accept  json
// @Produce  json
// @Param name query string true "标签名称"
// @Success 200
// @Router /api/v1/tags [post]
func CreateTag(c *gin.Context) {

	filter := make(map[string]interface{})
	filter["name"] = c.Query("name")

	tag := service.CreatedTag(filter)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "CreateTag",
		"data": tag,
	})
}

// @Tags 标签管理
// @Summary 修改标签
// @Description 根据传入的参数修改标签
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param id path int true "标签ID"
// @Param name formData string true "标签名称"
// @Success 200
// @Router /api/v1/tags/{id} [put]
func UpdateTag(c *gin.Context) {
	filter := make(map[string]string)
	filter["id"] = c.Param("id")
	filter["name"] = c.PostForm("name")
	tag := service.UpdateTag(filter)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "UpdateTag",
		"data": tag,
	})
}

// @Tags 标签管理
// @Summary 删除标签
// @Description 根据传入的参数删除标签
// @Accept  json
// @Produce  json
// @Param id path int false "标签ID"
// @Success 200
// @Router /api/v1/tags/{id} [delete]
func DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	tag := service.DeleteTag(id)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "DeleteTag",
		"data": tag,
	})
}
