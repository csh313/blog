package advert_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"log"
	"server/global"
	"server/models"
	"server/models/res"
	"server/service/pageService/common"
)

type AdvertApi struct{}

// AdvertListView 广告列表
// @Tags 广告管理
// @Summary 广告列表
// @Description 广告列表
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/advert [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}
func (AdvertApi) ShowAdvert(c *gin.Context) {
	var cr models.PageInfo
	var Advert models.AdvertModel
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, cr, c)
		return
	}
	list, count, err := common.PageList(cr, Advert)
	if err != nil {
		res.FailWithError(err, list, c)
		return
	}
	res.OkWithList(list, count, c)

}

// AdvertCreateView 添加广告
// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @Param data body models.AdvertModel    true  "表示多个参数"
// @Router /api/advert [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (AdvertApi) AdvertCreate(c *gin.Context) {

	var cr models.AdvertModel
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, cr, c)
		//res.FailWithError(err, cr, c)
		return
	}

	//重复判断
	if Row := global.DB.Take(&models.AdvertModel{}, "title=?", cr.Title).RowsAffected; Row != 0 {
		res.FailWithMessage("数据库中已存在", c)
		return
	}
	if err = global.DB.Create(&cr).Error; err != nil {
		res.FailWithMessage("添加广告失败", c)
		return
	}
	res.OkWithMessage("添加广告成功", c)
}

// AdvertRemoveView 批量删除广告
// @Tags 广告管理
// @Summary 批量删除广告
// @Description 批量删除广告
// @Param data body models.RemoveRequest    true  "广告id列表"
// @Router /api/advert [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (AdvertApi) AdvertDelete(c *gin.Context) {

	var ids models.RemoveRequest
	if err := c.ShouldBind(&ids); err != nil {
		res.FailWithError(err, ids, c)
		return
	}
	var advertList []models.AdvertModel
	log.Println("=============", ids)

	if count := global.DB.Find(&advertList, ids.IDList).RowsAffected; count == 0 {
		res.FailWithMessage("广告不存在", c)
		return
	}

	//if err := global.DB.Delete(&advertList).Error;
	if err := global.DB.Debug().Where("id IN ?", ids.IDList).Delete(&models.AdvertModel{}).Error; err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("删除广告成功", c)

}

// AdvertUpdateView 更新广告
// @Tags 广告管理
// @Summary 更新广告
// @Description 更新广告
// @Param data body models.AdvertModel    true  "广告的一些参数"
// @Param id path int true "id"
// @Router /api/advert/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (AdvertApi) AdvertUpdate(c *gin.Context) {
	id := c.Param("id")
	var cr models.AdvertModel

	if err := global.DB.Model(cr).Where("id=?", id).Take(&cr).Error; err != nil {
		res.FailWithError(err, cr, c)
		return
	}

	err := c.ShouldBindJSON(&cr)
	log.Println(cr)
	if err != nil {
		res.FailWithError(err, cr, c)
		return
	}

	maps := structs.Map(&cr)

	//Updates(map[string]any{
	//	"title":   cr.Title,
	//	"href":    cr.Href,
	//	"images":  cr.Images,
	//	"is_show": cr.IsShow,
	//}
	if err = global.DB.Model(&models.AdvertModel{}).Where("id=?", id).Updates(maps).Error; err != nil {
		res.FailWithError(err, cr, c)
		return
	}
	res.OkWithMessage("修改数据成功", c)

}
