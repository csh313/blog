package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"server/global"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routerGroup := router.Group("/api")
	SettingsRouter(routerGroup)
	ImagesRouter(routerGroup)
	AdvertRouter(routerGroup)
	UserRouter(routerGroup)
	TagRouter(routerGroup)
	MessageRouter(routerGroup)
	ArticleRouter(routerGroup)
	DiggRouter(routerGroup)
	return router
}
