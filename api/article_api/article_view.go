package article_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/service/pageService/common"
	"server/service/redis_service"
	"strconv"
)

type ArticleListRequest struct {
	models.PageInfo
	Tag string `form:"tag"` //Query参数 用form来接收
}

// ArticleList 文章列表
func (ArticleApi) ArticleList(c *gin.Context) {
	var cr ArticleListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var articleList []models.ArticleModel
	list, count, err := common.PageList(cr.PageInfo, articleList)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	count = global.DB.Take(&list, "tag", cr.Tag).RowsAffected

	res.OkWithList(list, count, c)
}

// 文章详情
func (ArticleApi) ArticleDetail(c *gin.Context) {
	var cr string
	if err := c.ShouldBindUri(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//查看文章详情时，浏览量增加
	if err := redis_service.Look(cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var article models.ArticleModel
	id, _ := strconv.Atoi(cr)
	if err := global.DB.Take(&article, "id=?", id).Error; err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	res.OkWithData(article, c)
}
