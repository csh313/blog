package user_api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"server/global"
	"server/models"
	"server/models/res"
	"server/plugins/email"
	"server/utils"
	"server/utils/pwd"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱格式不正确"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

func (UserApi) UserBindEmail(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(jwt.MapClaims)
	//用户绑定邮箱，第一次输入是邮箱
	//后台给邮箱发验证码
	var cr BindEmailRequest
	err := c.ShouldBind(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	session := sessions.Default(c)
	if cr.Code == nil {
		//第一次，后台发验证码，4位
		code := utils.Code(4)
		//生成验证码，将验证码存入session
		session.Set("valid_code", code)
		session.Save()
		err := email.NewCode().Send(cr.Email, "你的验证码是："+code)
		if err != nil {
			res.FailWithError(err, &cr, c)
			return
		}
		res.OkWithMessage("验证码已发送，请查收", c)
		return
	}

	code := session.Get("valid_code")
	//第二次用户输入邮箱、验证码、密码
	//校验验证码
	if code != *cr.Code {
		res.FailWithMessage("验证码错误", c)
		return
	}

	//修改用户邮箱  可能以前没有邮箱
	var user models.UserModel
	if err = global.DB.Take(&user, claims["UserID"]).Error; err != nil {
		res.OkWithMessage("用户不存在", c)
		return
	}

	hashPwd := pwd.HashPwd(cr.Password)
	global.DB.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashPwd,
	})
}
