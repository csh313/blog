package core

import (
	"github.com/go-redis/redis"
	"log"
	"server/global"
)

func InitRedis() {
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.IP + ":" + global.Config.Redis.Port,
		Password: global.Config.Redis.Password,
		DB:       0,
		PoolSize: global.Config.Redis.PoolSize,
	})
	if _, err := global.RedisClient.Ping().Result(); err != nil {
		log.Fatal("不能连接到redis服务器:", err)
	}
}
