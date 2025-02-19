package redis_service

import (
	"server/global"
	"strconv"
)

const lookPrefix = "look"

// 浏览一篇文章
func Look(id string) error {
	num, _ := global.RedisClient.HGet(lookPrefix, id).Int()
	num++
	err := global.RedisClient.HSet(lookPrefix, id, num).Err()
	return err
}

// 获取某一篇文章浏览量
func GetLook(id string) int {
	num, _ := global.RedisClient.HGet(lookPrefix, id).Int()
	return num
}

// 取出浏览量数据
func GetLookInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := global.RedisClient.HGetAll(lookPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo
}

// 清除浏览数据
func LookClear() {
	global.RedisClient.Del(lookPrefix)
}
