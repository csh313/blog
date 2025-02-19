package settings_api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"server/config"
	"server/core"
	"server/global"
	"server/models/res"
)

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	name := c.Param("name")
	switch name {
	case "site":
		var site config.SiteInfo
		err := c.ShouldBind(&site)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.SiteInfo = site
	case "email":
		var si config.Email
		//将请求体中json数据绑定到si结构体中
		err := c.ShouldBindJSON(&si)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		//配置信息修改
		global.Config.Email = si
	case "jwt":
		var si config.Jwt
		//将请求体中json数据绑定到si结构体中
		err := c.ShouldBindJSON(&si)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		//配置信息修改
		global.Config.Jwt = si
	case "qiniu":
		var si config.QiNiu
		//将请求体中json数据绑定到si结构体中
		err := c.ShouldBindJSON(&si)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		//配置信息修改
		global.Config.QiNiu = si
	case "qq":
		var si config.QQ
		//将请求体中json数据绑定到si结构体中
		err := c.ShouldBindJSON(&si)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		//配置信息修改
		global.Config.QQ = si
	case "upload":
		var si config.Upload
		//将请求体中json数据绑定到si结构体中
		err := c.ShouldBindJSON(&si)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		//配置信息修改
		global.Config.Upload = si
	default:
		res.FailWithMessage("没有对应的配置信息", c)
		return
	}
	//配置信息写入到yaml中(储存)
	err := core.SetYaml()
	if err != nil {
		//服务端查看
		logrus.Error(err)
		//返回客户端
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("修改成功", c)
	//var cr config.SiteInfo
	//err := c.ShouldBind(&cr)
	//if err != nil {
	//	res.FailWithCode(res.ArgumentError, c)
	//	return
	//}
	//global.Config.SiteInfo = cr
	//err = core.SetYaml()
	//if err != nil {
	//	global.Log.Error(err)
	//	res.FailWithMessage(err.Error(), c)
	//	return
	//}
	//res.OkWithData(cr, c)
}
