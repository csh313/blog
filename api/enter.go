package api

import (
	"server/api/advert_api"

	"server/api/images_api"
	"server/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
}

//var ApiGroupApp = new(ApiGroup)
