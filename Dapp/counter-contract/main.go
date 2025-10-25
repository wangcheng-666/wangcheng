package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"counter-contract/go-contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Println("=== 简单的以太坊合约交互程序 ===")
	fmt.Println("提示：这是一个示例程序，需要替换RPC URL和私钥才能运行")

	// 配置信息（需要修改）
	rpcURL := "https://sepolia.infura.io/v3/YOUR_PROJECT_ID" // 替换为您的Infura项目ID
	privateKeyHex := "YOUR_PRIVATE_KEY" // 替换为您的私钥

	// 步骤1: 连接到Sepolia测试网
	fmt.Println("\n步骤1: 连接到Sepolia测试网...")
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("无法连接到以太坊节点: %v", err)
	}
	defer client.Close()
	fmt.Println("✓ 已连接到Sepolia测试网")

	// 步骤2: 获取链ID
	fmt.Println("\n步骤2: 获取链ID...")
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("获取链ID失败: %v", err)
	}
	fmt.Printf("✓ 链ID: %v\n", chainID)

	// 步骤3: 解析私钥
	fmt.Println("\n步骤3: 解析私钥...")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("解析私钥失败: %v\n请检查私钥格式是否正确，确保没有0x前缀", err)
	}
	fmt.Println("✓ 私钥解析成功")

	// 步骤4: 创建交易授权
	fmt.Println("\n步骤4: 创建交易授权...")
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("创建交易授权失败: %v", err)
	}
	auth.GasLimit = uint64(300000) // 设置Gas上限
	fmt.Println("✓ 交易授权创建成功")

	// 步骤5: 部署合约
	fmt.Println("\n步骤5: 开始部署合约...")
	contractAddr, tx, instance, err := counter.DeployCounter(auth, client)
	if err != nil {
		log.Fatalf("部署合约失败: %v\n请确保您的账户在Sepolia测试网有足够的测试ETH", err)
	}

	fmt.Printf("✓ 合约部署交易哈希: %s\n", tx.Hash().Hex())
	fmt.Printf("✓ 合约地址: %s\n", contractAddr.Hex())
	fmt.Println("等待交易确认...")
	time.Sleep(30 * time.Second) // 等待交易确认

	// 步骤6: 调用合约的increase方法
	fmt.Println("\n步骤6: 调用increase方法增加计数...")
	tx, err = instance.Increase(auth)
	if err != nil {
		log.Fatalf("调用increase方法失败: %v", err)
	}
	fmt.Printf("✓ 交易哈希: %s\n", tx.Hash().Hex())
	fmt.Println("等待交易确认...")
	time.Sleep(30 * time.Second) // 等待交易确认

	// 步骤7: 读取计数
	fmt.Println("\n步骤7: 读取当前计数...")
	count, err := instance.GetCount(nil)
	if err != nil {
		log.Fatalf("获取计数失败: %v", err)
	}
	fmt.Printf("✓ 当前计数: %d\n", count)

	fmt.Println("\n🎉 程序执行完成！")
	fmt.Println("\n提示：")
	fmt.Println("1. 您可以在Etherscan上查看交易详情：https://sepolia.etherscan.io/")
	fmt.Println("2. 如果需要与已部署的合约交互，可以修改代码使用合约地址进行连接")
}
