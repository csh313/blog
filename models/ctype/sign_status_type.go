package ctype

import "encoding/json"

// 注册来源类型
type SignStatus int

// 注册来源数据
const (
	SignQQ    SignStatus = 1 // 注册来源QQ
	SignGitee SignStatus = 2 // 注册来源Gitee
	SignEmail SignStatus = 3 // 注册来源邮箱
)

// 注册来源序列化
func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// 注册来源匹配
func (s SignStatus) String() string {
	switch s {
	case SignQQ:
		return "QQ"
	case SignGitee:
		return "Gitee"
	case SignEmail:
		return "Email"
	default:
		return "未知来源"
	}
}
