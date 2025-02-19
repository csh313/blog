package tag_api

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/service/pageService/common"
)

// 广告上传请求参数
type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标签名"  structs:"title"` //标签名
}

// TagCreateView 添加标签
// @Tags 标签管理
// @Summary 添加标签
// @Description 添加标签
// @Param data body TagRequest    true  "表示多个参数"
// @Router /api/tags [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (TagApi) TagCreateView(c *gin.Context) {

	var req TagRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}

	//重复判断，是否已添加
	var tag models.TagModel
	err = global.DB.Take(&tag, "title = ?", req.Title).Error
	if err == nil {
		res.FailWithMessage("该标签已存在", c)
		return
	}

	err = global.DB.Create(&models.TagModel{
		Title: req.Title,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.OkWithMessage("添加标签失败", c)
		return
	}
	res.OkWithMessage("添加标签成功", c)

}

// TagRemoveView 批量删除标签
// @Tags 标签管理
// @Summary 批量删除标签
// @Description 批量删除标签
// @Param data body models.RemoveRequest    true  "标签id列表"
// @Router /api/tags [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var tagList []models.TagModel
	count := global.DB.Find(&tagList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("标签不存在", c)
		return
	}

	// todo删除标签是标签关联有文章该怎么办？

	global.DB.Delete(&tagList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个标签", count), c)

}

// TagUpdateView 更新标签
// @Tags 标签管理
// @Summary 更新标签
// @Description 更新标签
// @Param data body TagRequest    true  "标签的一些参数"
// @Param id path int true "id"
// @Router /api/tags/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (TagApi) TagUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var tag models.TagModel
	err = global.DB.First(&tag, id).Error
	if err != nil {
		res.FailWithMessage("标签不存在", c)
		return
	}
	// 结构体转map的第三方包
	// 用map原因，让布尔值可以正常修改
	// 删除binding:"required" msg:"请选择是否显示"
	maps := structs.Map(&cr)
	fmt.Println(maps)
	err = global.DB.Model(&tag).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改标签失败", c)
		return
	}

	res.OkWithMessage("修改标签成功", c)
}

// TagListView 标签列表
// @Tags 标签管理
// @Summary 标签列表
// @Description 标签列表
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/tags [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.TagModel]}
func (TagApi) TagListView(c *gin.Context) {

	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var tagList models.TagModel
	//使用封装的列表分页查询服务
	list, count, err := common.PageList(page, tagList)
	if err != nil {
		res.FailWithMessage("标签列表获取失败", c)
		return
	}

	//需要显示这个标签下关联文章的数量

	res.OkWithList(list, count, c)
}
