package user_api

import (
	"fmt"
	"server/global"
	"server/models/ctype"
	"server/models/res"
	"server/service/userService"

	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	NickName string     `json:"nick_name" binding:"required" msg:"请输入昵称"`
	UserName string     `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string     `json:"password" binding:"required" msg:"请输入密码"`
	Role     ctype.Role `json:"role"     binding:"required" msg:"请选择权限"`
}

// UserCreateView 创建用户
// @Tags 用户管理
// @Summary 创建用户
// @Description 创建用户
// @Param data body UserCreateRequest    true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/users [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserCreateView(c *gin.Context) {
	var req UserCreateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	err = user_service.UserService{}.CreateUser(req.NickName, req.UserName, req.Password, req.Role, "", c.ClientIP())
	// 创建失败
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	// 创建成功
	res.OkWithMessage(fmt.Sprintf("用户%s创建成功!", req.UserName), c)
}
