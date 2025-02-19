package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"server/global"
	"server/models/res"
	"time"
)

func (UserApi) Logout(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(jwt.MapClaims)

	token := c.Request.Header.Get("Authorization")

	//计算过期时间
	exp := claims["exp"]
	now := time.Now()

	diff := exp.(time.Time).Sub(now)

	if err := global.RedisClient.Set(fmt.Sprintf("logout_%s", token), token, diff).Err(); err != nil {
		global.Log.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}
	res.OkWithMessage("注销成功", c)
}
