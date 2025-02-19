package redis_service

import (
	"server/global"
	"strconv"
)

const diggerPrefix = "digg"

// 点赞一篇文章
func Digg(id string) error {
	nums, _ := global.RedisClient.HGet(diggerPrefix, id).Int()
	nums++
	//更新点赞数
	err := global.RedisClient.HSet(diggerPrefix, id, nums).Err()
	return err
}

// 获取文章的点赞数
func GetDigg(id string) int {
	nums, _ := global.RedisClient.HGet(diggerPrefix, id).Int()
	return nums
}

// 取出点赞数据
func GetDiggInfo() map[string]int {
	var DiggInfo = make(map[string]int)
	maps := global.RedisClient.HGetAll(diggerPrefix).Val()
	for k, v := range maps {
		num, _ := strconv.Atoi(v)
		DiggInfo[k] = num
	}
	return DiggInfo
}

// 清空点赞数
func DiggClearAll() {
	global.RedisClient.Del(diggerPrefix)
}

// 清除某几个id的点赞数
func DiggClear(list []string) {
	global.RedisClient.HDel(diggerPrefix, list...)
}
