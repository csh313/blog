package routers

import (
	"github.com/gin-gonic/gin"
	"server/api/tag_api"
)

func TagRouter(router *gin.RouterGroup) {
	TagApi := tag_api.TagApi{}

	router.POST("/tag", TagApi.TagCreateView)
	router.GET("/tags", TagApi.TagListView)
	router.PUT("/tag/:id", TagApi.TagUpdateView)
	router.DELETE("/tags", TagApi.TagRemoveView)
}
