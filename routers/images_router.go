package routers

import (
	"github.com/gin-gonic/gin"
	"server/api/images_api"
)

func ImagesRouter(router *gin.RouterGroup) {
	ImagesApi := images_api.ImagesApi{}
	router.POST("/image", ImagesApi.ImageUploadView)
	router.GET("/image", ImagesApi.ImageListView)
	router.DELETE("/image", ImagesApi.ImageRemoveView)
	router.PUT("/image", ImagesApi.ImageUpdateView)
}
