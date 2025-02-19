package digger_api

import (
	"github.com/gin-gonic/gin"
	"server/models/res"
	"server/service/redis_service"
)

type DiggRequest struct {
	ID string `uri:"id" json:"id" form:"id"`
}

func (DiggerApi) DiggArticleView(c *gin.Context) {
	var cr DiggRequest
	if err := c.ShouldBindUri(&cr); err != nil {
		res.FailWithError(err, cr, c)
		return
	}
	//给文章点赞
	//todo redis点赞和数据库中的点赞，只实现了一个接口，没有其他作用
	if err := redis_service.Digg(cr.ID); err != nil {
		res.FailWithMessage("点赞失败", c)
		return
	}

	res.OkWithMessage("点赞成功", c)
}
