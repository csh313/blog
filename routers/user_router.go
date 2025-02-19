package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"server/api/user_api"
	"server/middleware"
)

func UserRouter(router *gin.RouterGroup) {
	var store = cookie.NewStore([]byte("SESSION"))
	router.Use(sessions.Sessions("SessionID", store))
	UserApi := user_api.UserApi{}
	router.POST("/login", UserApi.Login)
	router.POST("/register", UserApi.UserCreateView)

	//只有登录的用户才能调用用户信息列表
	router.GET("/users", middleware.JwtAuth(), UserApi.UserList)
	router.PUT("/user_pwd", middleware.JwtAuth(), UserApi.UserUpdatePassword)
	router.POST("/user_bind_email", middleware.JwtAuth(), UserApi.UserBindEmail)
	router.POST("/logout", middleware.JwtAuth(), UserApi.Logout)

	//管理员修改用户权限
	router.PUT("/user_role", middleware.JwtAdmin(), UserApi.UserRemoveView)
}
