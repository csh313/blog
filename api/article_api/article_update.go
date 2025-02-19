package article_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/ctype"
	"server/models/res"
)

type ArticleUpdateRequest struct {
	Title    string   `json:"title" `    // 文章标题
	Abstract string   `json:"abstract" ` // 文章简介
	Content  string   `json:"content" `  // 文章内容
	Category string   `json:"category"`  // 文章分类
	Source   string   `json:"source"`    // 文章来源
	Link     string   `json:"link"`      // 原文链接
	BannerID uint     `json:"banner_id"` // 文章封面id
	Tags     []string `json:"tags"`      // 文章标签
	ID       string   `json:"id" `
}

func (ArticleApi) ArticleUpdate(c *gin.Context) {
	var cr ArticleUpdateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, cr, c)
		return
	}

	var article models.ArticleModel

	//查看数据库中是否包含该文章
	if err := global.DB.Take(&article, "id=?", cr.ID).Error; err != nil {
		res.FailWithMessage("数据库中不存在该文章", c)
		return
	}

	maps := structs.Map(cr)
	var DataMap = map[string]any{}
	for key, v := range maps {
		switch val := v.(type) {
		case string:
			if val == "" {
				continue
			}
		case uint:
			if val == 0 {
				continue
			}
		case int:
			if val == 0 {
				continue
			}
		case ctype.Array:
			if len(val) == 0 {
				continue
			}
		}
		DataMap[key] = v
	}

	if err := global.DB.Model(&article).Updates(DataMap).Error; err != nil {
		res.FailWithError(err, cr, c)
		return
	}
	res.OkWithMessage("修改文章成功", c)

}
