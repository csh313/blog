package global

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"server/config"
)

var (
	Config      *config.Config
	DB          *gorm.DB
	Log         *logrus.Logger
	MysqlLog    logger.Interface
	RedisClient *redis.Client
)
