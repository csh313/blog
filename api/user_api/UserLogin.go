package user_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/utils/pwd"
)

type EmailLoginRequest struct {
	Username string `json:"username" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (UserApi) Login(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var userModel models.UserModel
	err = global.DB.Where("user_name=?", cr.Username).Find(&userModel).Error
	if err != nil {
		global.Log.Warn("用户名不存在")
		res.FailWithMessage("用户名不存在", c)
		return
	}
	//校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")
		res.FailWithMessage("用户名密码错误", c)
		return
	}
	token, err := pwd.GenerateJWT(pwd.JwtPayLoad{
		Username: userModel.UserName,
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		res.FailWithMessage("token生成失败", c)
		return
	}
	res.OkWithData(token, c)
}
