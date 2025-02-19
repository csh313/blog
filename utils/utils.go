package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//func GetValidMsg(err error, obj any) string {
//	return "FailWithError"
//}

// 判断字符串是否存在于list中
func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

// MD5加密函数,在项目中用于加密图片文件信息
func MD5(bytedata []byte) string {
	m := md5.New()
	m.Write(bytedata)
	//将计算得到的哈希值转换为16进制字符串
	str := hex.EncodeToString(m.Sum(nil))
	return str
}
