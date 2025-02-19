package config

import "fmt"

type Redis struct {
	IP       string `json:"ip" yaml:"ip"`               // redis地址
	Port     string `json:"port" yaml:"port"`           // redis端口
	Password string `json:"password" yaml:"password"`   // redis密码
	PoolSize int    `json:"pool_size" yaml:"pool_size"` // 最大连接数
}

// 拼接获取addr
func (r *Redis) Addr() string {
	return fmt.Sprintf("%s:%d", r.IP, r.Port)
}
