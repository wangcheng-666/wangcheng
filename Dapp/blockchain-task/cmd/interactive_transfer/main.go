package main

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"blockchain-task/internal/client"
	"blockchain-task/internal/config"
	"blockchain-task/internal/transaction"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 检查私钥配置
	if cfg.PrivateKey == "" {
		log.Fatalf("私钥未设置")
	}

	// 检查接收地址配置
	if cfg.RecipientAddr == "" {
		log.Fatalf("接收地址未设置")
	}

	// 解析私钥
	privateKey, err := transaction.ParsePrivateKey(cfg.PrivateKey)
	if err != nil {
		log.Fatalf("解析私钥失败: %v", err)
	}

	// 创建以太坊客户端
	ethClient, err := client.NewEthClient(cfg)
	if err != nil {
		log.Fatalf("创建以太坊客户端失败: %v", err)
	}
	defer ethClient.Close()

	// 设置转账金额为以太坊最低额度：1 wei
	amount := big.NewInt(1) // 1 wei = 0.000000000000000001 ETH
	recipientAddr := common.HexToAddress(cfg.RecipientAddr)

	fmt.Println("=== 最低额度测试交易 ===")
	fmt.Printf("转账金额: 1 wei (0.000000000000000001 ETH)\n")
	fmt.Printf("接收地址: %s\n", recipientAddr.Hex())
	fmt.Println("3秒后继续...")
	time.Sleep(3 * time.Second)

	// 签名并发送交易
	txHash, err := transaction.SignAndSendTransaction(ethClient, privateKey, recipientAddr, amount)
	if err != nil {
		log.Fatalf("发送交易失败: %v", err)
	}

	fmt.Printf("\n最低额度测试交易已发送！\n交易哈希: %s\n", txHash.Hex())
	fmt.Printf("查看交易详情: https://sepolia.etherscan.io/tx/%s\n\n", txHash.Hex())
	fmt.Println("重要说明：")
	fmt.Println("1. 由于金额极小（1 wei），在大多数钱包界面中可能看不到余额变化")
	fmt.Println("2. 请通过区块浏览器确认交易是否成功执行")
	fmt.Println("3. 如果需要实际看到余额变化，建议使用较大金额（如0.001 ETH）")
}
