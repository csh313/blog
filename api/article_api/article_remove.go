package article_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/service/redis_service"
)

type ArticleRemoveRequest struct {
	IDList []string `json:"id_list"`
}

func (ArticleApi) ArticleRemove(c *gin.Context) {
	var cr ArticleRemoveRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, cr, c)
		return
	}
	if err := global.DB.Delete(&models.ArticleModel{}, cr.IDList).Error; err != nil {
		res.FailWithMessage("删除失败", c)
		return
	}
	//删除点赞数
	redis_service.DiggClear(cr.IDList)

	res.OkWithMessage("删除成功", c)
}
