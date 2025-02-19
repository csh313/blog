package config

// 配置结构体
type AliyunOSS struct {
	Enable          bool   `json:"enable" yaml:"enable"`
	BucketName      string `yaml:"bucket_name" json:"bucket_name"`
	Endpoint        string `yaml:"endpoint" json:"endpoint"`
	AccessKeyID     string `yaml:"access_key_id" json:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret" json:"access_key_secret"`
	Region          string `yaml:"region" json:"region"`
}
