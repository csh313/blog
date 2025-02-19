package user_api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"server/models"
	"server/models/ctype"
	"server/models/res"
	"server/service/pageService/common"
)

func (UserApi) UserList(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(jwt.MapClaims)
	role := claims["role"]

	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := common.PageList(page, models.UserModel{})
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	var users []models.UserModel
	for _, user := range list {
		if role != ctype.PermissionAdmin {
			user.UserName = " "
		}
		users = append(users, user)
	}
	res.OkWithList(users, count, c)

}
