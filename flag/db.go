package flag

import (
	"server/global"
	"server/models"

	"github.com/sirupsen/logrus"
)

func Makemigrations() {
	var err error
	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "CollectsModels", &models.MenuBannerModel{})
	//生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},
			&models.TagModel{},
			&models.MessageModel{},
			&models.AdvertModel{},
			&models.UserModel{},
			&models.CommentModel{},
			&models.ArticleModel{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			&models.LoginDataModel{},
			&models.FadeBackModel{},
		)
	if err != nil {
		logrus.Error("初始化数据库失败", err)
		return
	}
	logrus.Info("初始化数据库成功")
}
