package user_api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"server/global"
	"server/models"
	"server/models/res"
	"server/utils/pwd"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd"` // 旧密码
	Pwd    string `json:"pwd"`     // 新密码
}

func (UserApi) UserUpdatePassword(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(jwt.MapClaims)

	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var user models.UserModel
	if err := global.DB.Take(&user, claims["UserID"]).Error; err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	//判断旧密码是否正确
	if !pwd.CheckPwd(user.Password, cr.OldPwd) {
		res.FailWithMessage("原密码错误", c)
		return
	}
	//判断新密码是否符合规范
	if cr.Pwd == "" {
		res.FailWithMessage("新密码不能为空", c)
		return
	}
	if cr.Pwd == cr.OldPwd {
		res.FailWithMessage("新密码不能与原密码相同", c)
		return
	}

	//加密新密码
	hashPwd := pwd.HashPwd(cr.Pwd)
	if err := global.DB.Model(&user).Update("password", hashPwd).Error; err != nil {
		res.FailWithError(err, "密码修改失败", c)
		return
	}
	res.OkWithMessage("密码修改成功", c)

}
