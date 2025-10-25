package transaction

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// SignAndSendTransaction 签名并发送以太币转账交易
func SignAndSendTransaction(client *ethclient.Client, privateKey *ecdsa.PrivateKey, to common.Address, amount *big.Int) (common.Hash, error) {
	// 获取发送者地址
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// 获取nonce值 ，这里的context就是控制noce的生命周期，这样写会一直运行直到程序终止
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return common.Hash{}, fmt.Errorf("获取nonce失败: %w", err)
	}

	// 设置燃气价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return common.Hash{}, fmt.Errorf("获取燃气价格失败: %w", err)
	}

	// 设置燃气上限
	gasLimit := uint64(21000) // 标准转账交易的燃气上限

	// 创建交易对象
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, nil)

	// 获取链ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return common.Hash{}, fmt.Errorf("获取链ID失败: %w", err)
	}

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return common.Hash{}, fmt.Errorf("签名交易失败: %w\n请检查您的私钥是否正确。", err)
	}

	// 发送交易
	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		return common.Hash{}, fmt.Errorf("发送交易失败: %w\n可能的原因：余额不足、nonce错误或网络问题。", err)
	}

	// 返回交易哈希
	txHash := signedTx.Hash()
	// 移除log.Printf，让main.go来处理输出

	return txHash, nil
}

// ParsePrivateKey 将十六进制格式的私钥字符串解析为ecdsa.PrivateKey
func ParsePrivateKey(privateKeyHex string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("解析私钥失败: %w\n请确保私钥格式正确（不包含0x前缀的64位十六进制字符串）。", err)
	}

	return privateKey, nil
}
