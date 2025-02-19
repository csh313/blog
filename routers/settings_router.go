package routers

import (
	"github.com/gin-gonic/gin"
	"server/api/settings_api"
)

func SettingsRouter(router *gin.RouterGroup) {
	SettingsApi := settings_api.SettingsApi{}
	router.GET("/info/:name", SettingsApi.SettingsInfoView)
	router.PUT("/info/:name", SettingsApi.SettingsInfoUpdateView)

}
