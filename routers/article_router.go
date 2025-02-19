package routers

import (
	"github.com/gin-gonic/gin"
	"server/api/article_api"
	"server/middleware"
)

func ArticleRouter(router *gin.RouterGroup) {
	ArticleApi := article_api.ArticleApi{}
	router.POST("/article", middleware.JwtAuth(), ArticleApi.ArticleCreateView)
	router.GET("/articles", ArticleApi.ArticleList)
	router.DELETE("/articles", ArticleApi.ArticleRemove)

	router.GET("/article/:id", ArticleApi.ArticleDetail)
	router.PUT("/article", ArticleApi.ArticleUpdate)

}
