package routers

import (
	"github.com/gin-gonic/gin"
	"server/api/digger_api"
)

func DiggRouter(router *gin.RouterGroup) {
	DiggerApi := digger_api.DiggerApi{}
	router.POST("/digg/article", DiggerApi.DiggArticleView)
}
