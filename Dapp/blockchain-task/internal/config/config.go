package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config 存储应用程序的配置信息
type Config struct {
	InfuraAPIKey  string
	PrivateKey    string
	RecipientAddr string
}

// LoadConfig 从.env文件和环境变量中加载配置
func LoadConfig() (*Config, error) {
	// 尝试加载.env文件
	_ = godotenv.Load()

	// 从环境变量中获取配置
	infuraAPIKey := os.Getenv("INFURA_API_KEY")
	privateKey := os.Getenv("PRIVATE_KEY")
	recipientAddr := os.Getenv("RECIPIENT_ADDRESS")

	// 验证配置
	if infuraAPIKey == "" {
		return nil, fmt.Errorf("INFURA_API_KEY")
	}

	// 创建并返回配置实例
	return &Config{
			InfuraAPIKey:  infuraAPIKey,
			PrivateKey:    privateKey,
			RecipientAddr: recipientAddr,
		},
		nil
}
