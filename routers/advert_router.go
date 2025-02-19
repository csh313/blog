package routers

import (
	"github.com/gin-gonic/gin"
	"server/api/advert_api"
)

func AdvertRouter(router *gin.RouterGroup) {
	AdvertApi := advert_api.AdvertApi{}
	router.GET("/advert", AdvertApi.ShowAdvert)
	router.POST("/advert", AdvertApi.AdvertCreate)
	router.DELETE("/advert", AdvertApi.AdvertDelete)
	router.PUT("/advert/:id", AdvertApi.AdvertUpdate)
}
