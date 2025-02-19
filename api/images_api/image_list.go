package images_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

func (ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBind(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var imageList []models.BannerModel

	//count:=global.DB.Find(&imageList).RowsAffected
	offset := (cr.Page - 1) * cr.Limit
	if offset < 0 {
		offset = 0
	}
	if cr.Sort == "" {
		cr.Sort = "created_at desc" //按时间降序
	}
	if cr.Limit > 0 {
		global.DB.Debug().Limit(cr.Limit).Offset(offset).Order(cr.Sort).Find(&imageList)
	} else {
		global.DB.Debug().Offset(offset).Order(cr.Sort).Find(&imageList)
	}
	//global.DB.Find(&imageList)
	res.OkWithData(imageList, c)
}

func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	var imageList []models.BannerModel
	err := c.ShouldBind(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	if count := global.DB.Find(&imageList, cr.IDList).RowsAffected; count == 0 {
		res.FailWithMessage("文件不存在", c)
	}

	if err = global.DB.Debug().Delete(&imageList).Error; err != nil {
		res.FailWithMessage("删除图片失败", c)
	}
	res.OkWithMessage("删除图片成功", c)

}
