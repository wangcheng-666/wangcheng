package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
	"blockchain-task/internal/config"
	"blockchain-task/internal/client"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 解析私钥
	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatalf("解析私钥失败: %v", err)
	}

	// 获取发送者地址
	senderAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	recipientAddress := common.HexToAddress(cfg.RecipientAddr)

	// 创建以太坊客户端
	ethClient, err := client.NewEthClient(cfg)
	if err != nil {
		log.Fatalf("创建以太坊客户端失败: %v", err)
	}
	defer ethClient.Close()

	fmt.Println("==================== 诊断信息 ====================")
	fmt.Printf("发送地址: %s\n", senderAddress.Hex())
	fmt.Printf("接收地址: %s\n", recipientAddress.Hex())

	// 检查发送者余额
	senderBalance, err := ethClient.BalanceAt(context.Background(), senderAddress, nil)
	if err != nil {
		log.Fatalf("获取发送者余额失败: %v", err)
	}
	fmt.Printf("发送者余额: %s Wei (约 %f ETH)\n", senderBalance.String(), float64(senderBalance.Int64())/1e18)

	// 检查接收者余额
	recipientBalance, err := ethClient.BalanceAt(context.Background(), recipientAddress, nil)
	if err != nil {
		log.Fatalf("获取接收者余额失败: %v", err)
	}
	fmt.Printf("接收者余额: %s Wei (约 %f ETH)\n", recipientBalance.String(), float64(recipientBalance.Int64())/1e18)

	// 检查发送地址的nonce
	nonce, err := ethClient.PendingNonceAt(context.Background(), senderAddress)
	if err != nil {
		log.Fatalf("获取nonce失败: %v", err)
	}
	fmt.Printf("发送地址当前nonce: %d\n", nonce)

	// 建议解决方案
	fmt.Println("\n==================== 问题分析 ====================")
	if senderBalance.Cmp(big.NewInt(0)) <= 0 {
		fmt.Println("问题: 发送地址余额为0或负数！")
		fmt.Println("解决方案: 请从Sepolia水龙头获取测试以太币:\n  - https://www.infura.io/faucet/sepolia\n  - https://sepoliafaucet.com/\n  - https://faucets.chainstack.com/sepolia-faucet/")
	} else {
		fmt.Println("发送地址有足够余额。")
		fmt.Println("\n可能的问题原因:")
		fmt.Println("1. 交易可能还在等待矿工确认 (通常需要几秒到几分钟)")
		fmt.Println("2. 之前发送的金额太小 (如果发送的是1 wei，几乎等于0)")
		fmt.Println("3. 请确认接收地址是否正确")
		fmt.Println("\n建议操作:")
		fmt.Println("1. 发送一个更明显的金额（例如 0.001 ETH）")
		fmt.Println("   命令: go run cmd/send_transaction/main.go 1000000000000000")
		fmt.Println("2. 在Sepolia区块浏览器上检查交易状态")
		fmt.Println("   https://sepolia.etherscan.io/address/" + recipientAddress.Hex())
	}
	fmt.Println("==================================================")
}