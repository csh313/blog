package middleware

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/ctype"
	"server/models/res"
	"server/utils/pwd"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未登录", c)
			c.Abort()
			return
		}
		claims, err := pwd.ParseJWT(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		//判断是否在redis里
		keys := global.RedisClient.Keys("logout_").Val()
		for _, key := range keys {
			if "logout_"+token == key {
				res.FailWithMessage("token已失效", c)
				c.Abort()
				return
			}
		}
		//登录的用户
		c.Set("claims", claims)
	}
}

// 管理员调用的接口
func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未登录", c)
			c.Abort()
		}
		claims, err := pwd.ParseJWT(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
		}

		//登录的用户
		role := claims["role"]
		if role != ctype.PermissionAdmin {
			res.FailWithMessage("权限错误", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
