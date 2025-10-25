package client

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"blockchain-task/internal/config"
)

// NewEthClient 创建一个连接到Sepolia测试网络的以太坊客户端
func NewEthClient(cfg *config.Config) (*ethclient.Client, error) {
	// 构建Infura的Sepolia测试网络URL
	url := fmt.Sprintf("https://sepolia.infura.io/v3/%s", cfg.InfuraAPIKey)

	// 连接到以太坊网络
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("无法连接到Sepolia测试网络: %w\n请确保您的Infura API Key正确，并且Infura服务可访问。", err)
	}

	return client, nil
}