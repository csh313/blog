package main

import (
	"server/core"
	_ "server/docs"
	"server/flag"
	"server/global"
	"server/plugins/AliOss"
	"server/routers"
	"strconv"
)

// 添加注释以描述 server 信息
// @title           BLOG server
// @version         1.0
// @description     博客server
// @host      localhost:8080
// @BasePath  /
func main() {
	core.InitConf()
	core.InitGorm()
	core.InitRedis()
	core.InitLogger()
	//命令行参数绑定迁移表结构函数
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		//控制迁移表结构后退出
		return
	}
	//fmt.Println(global.DB)

	router := routers.InitRouter()
	AliOss.InitAliOss()

	router.Run(":" + strconv.Itoa(global.Config.System.Port))

}
