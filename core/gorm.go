package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"server/global"
	"time"
)

func InitGorm() *gorm.DB {

	if global.Config.Mysql.Host == "" {
		//log.Printf("未配置mysql,取消gorm连接")
		global.Log.Error("未配置mysql,取消gorm连接")
		return nil
	}

	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}

	global.MysqlLog = logger.Default.LogMode(logger.Info)

	DSN := global.Config.Mysql.Dsn()
	var err error
	global.DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
		QueryFields: true, //打印sql
		Logger:      mysqlLogger,
	})

	if err != nil {
		global.Log.Error(fmt.Sprintf("[%s] mysql连接失败", DSN))
		//global.LOG.Error(fmt.Sprintf("[%s] mysql连接失败", dsn))
		//panic(err)
	}
	sqlDB, _ := global.DB.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大服用时间,不能超过mysql的wait_timeout
	return global.DB
}
