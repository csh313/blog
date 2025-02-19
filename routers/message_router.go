package routers

import (
	"github.com/gin-gonic/gin"
	"server/api/message_api"
	"server/middleware"
)

func MessageRouter(router *gin.RouterGroup) {
	MessageApi := message_api.MessageApi{}

	router.POST("/message", MessageApi.MessageCreate)
	router.GET("/messages", middleware.JwtAuth(), MessageApi.MessageList)
	router.GET("/auth/messages_all", middleware.JwtAdmin(), MessageApi.MessageListAll)
}
