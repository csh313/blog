package utils

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

// 生成随机验证码
func Code(length int) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%4v", rand.Intn(10000))
}

// GenerateRandomSecret 生成随机的 secret，长度为 n 字节，并返回 base64 编码后的字符串
func GenerateRandomSecret(n int) (string, error) {
	// 创建一个字节数组来保存随机数据
	secret := make([]byte, n)

	// 从 crypto/rand 生成随机字节，保证安全性
	_, err := rand.Read(secret)
	if err != nil {
		return "", err
	}

	// 将字节编码为 base64 字符串
	return base64.StdEncoding.EncodeToString(secret), nil
}
