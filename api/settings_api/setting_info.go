package settings_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/res"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	//获取url的参数
	//name := c.Param("name")

	var uri SettingsUri
	// 绑定uri参数
	err := c.ShouldBindUri(&uri)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	name := uri.Name
	fmt.Println("name:" + name)
	//name := c.Param("name")
	switch name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "upload":
		res.OkWithData(global.Config.Upload, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}
}
