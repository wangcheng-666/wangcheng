package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"blockchain-task/internal/client"
	"blockchain-task/internal/config"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 创建以太坊客户端
	ethClient, err := client.NewEthClient(cfg)
	if err != nil {
		log.Fatalf("创建以太坊客户端失败: %v", err)
	}
	defer ethClient.Close()

	// 确定要查询的区块号
	var blockNumber *big.Int

	// 检查命令行参数
	if len(os.Args) > 1 {
		// 尝试解析命令行参数为区块号
		num, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatalf("无效的区块号: %v", err)
		}
		blockNumber = big.NewInt(num)
	} else {
		// 获取最新区块号
		latestBlock, err := ethClient.BlockByNumber(context.Background(), nil)
		if err != nil {
			log.Fatalf("获取最新区块失败: %v", err)
		}
		blockNumber = latestBlock.Number()
		fmt.Printf("未指定区块号，查询最新区块 #%d\n\n", blockNumber)
	}

	// 查询指定区块
	block, err := ethClient.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatalf("查询区块 #%d 失败: %v\n请确认区块号是否存在。", blockNumber, err)
	}

	// 输出区块信息
	printBlockInfo(block)

	// 如果区块有交易，输出交易信息
	if len(block.Transactions()) > 0 {
		printTransactionInfo(block, ethClient)
	}
}

// printBlockInfo 打印区块的基本信息
func printBlockInfo(block *types.Block) {
	fmt.Println("==================== 区块信息 ====================")
	fmt.Printf("区块号: %d\n", block.Number())
	fmt.Printf("区块哈希: %s\n", block.Hash().Hex())
	fmt.Printf("父区块哈希: %s\n", block.ParentHash().Hex())
	fmt.Printf("时间戳: %v\n", block.Time())
	fmt.Printf("矿工地址: %s\n", block.Coinbase().Hex())
	fmt.Printf("难度值: %v\n", block.Difficulty().Uint64())
	fmt.Printf("燃气上限: %d\n", block.GasLimit())
	fmt.Printf("燃气使用量: %d\n", block.GasUsed())
	fmt.Printf("交易数量: %d\n", len(block.Transactions()))
	fmt.Printf("叔区块数量: %d\n", len(block.Uncles()))
	fmt.Printf("区块大小: %d bytes\n", block.Size())
	fmt.Printf("状态根哈希: %s\n", block.Root().Hex())
	fmt.Printf("交易根哈希: %s\n", block.TxHash().Hex())
	fmt.Printf("收据根哈希: %s\n", block.ReceiptHash().Hex())
	fmt.Println("==================================================\n")
}

// printTransactionInfo 打印区块中的交易信息
func printTransactionInfo(block *types.Block, client *ethclient.Client) {
	fmt.Printf("区块中共有 %d 笔交易:\n\n", len(block.Transactions()))

	// 为了简化，只打印第一笔交易的信息
	tx := block.Transactions()[0]
	fmt.Println("交易信息:")
	fmt.Printf("交易哈希: %s\n", tx.Hash().Hex())
	
	// 获取接收方地址
	toAddress := tx.To()
	if toAddress == nil {
		fmt.Println("接收方: 合约创建交易")
	} else {
		fmt.Printf("接收方: %s\n", toAddress.Hex())
	}
	
	// 打印其他交易信息
	fmt.Printf("转账金额: %s wei\n", tx.Value().String())
	fmt.Printf("Gas价格: %s wei\n", tx.GasPrice().String())
	fmt.Printf("Gas上限: %d\n", tx.Gas())
	fmt.Printf("Nonce: %d\n\n", tx.Nonce())
	fmt.Println("提示：在实际项目中，可以使用交易签名恢复发送方地址。\n")
}