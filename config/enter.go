package config

// 存储配置信息结构体
type Config struct {
	Mysql     MySQL     `yaml:"mysql"`
	System    System    `yaml:"system"`
	Logger    Logger    `yaml:"logger"`
	SiteInfo  SiteInfo  `yaml:"site_info"`
	Email     Email     `yaml:"email"`
	Jwt       Jwt       `yaml:"jwt"`
	QiNiu     QiNiu     `yaml:"qi_niu"`
	QQ        QQ        `yaml:"qq"`
	Upload    Upload    `yaml:"upload"`
	Redis     Redis     `yaml:"redis"`
	ES        ES        `yaml:"es"`
	AliyunOSS AliyunOSS `yaml:"ali_oss"`
}

type UpdateConfigRequest struct {
	SiteInfo  SiteInfo  `json:"siteinfo,omitempty"`
	Email     Email     `json:"email,omitempty"`
	Jwt       Jwt       `json:"jwt,omitempty"`
	QiNiu     QiNiu     `json:"qiniu,omitempty"`
	QQ        QQ        `json:"qq,omitempty"`
	Upload    Upload    `json:"upload,omitempty"`
	AliyunOSS AliyunOSS `json:"ali_oss,omitempty"`
}
