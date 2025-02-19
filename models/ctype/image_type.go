package ctype

import "encoding/json"

//枚举类型

// 图片存储类型，用于图片删除时判断删除方式
type ImageType int

const (
	Local       ImageType = 1 //本地存储
	QiNiu        ImageType = 2 //七牛云存储

)

// 角色json序列化
func (r ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// 角色匹配
func (r ImageType) String() string {
	switch r {
	case Local:
		return "本地存储"
	case QiNiu:
		return "七牛云存储"
	default:
		return "其他"
	}
}
