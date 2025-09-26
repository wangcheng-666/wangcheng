package utils

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// 🟡 将 map 改为首字母大写，但不直接暴露
// 使用 Getter/Setter 控制访问
var challenges = make(map[string]string)

// GetChallenge 存储并返回挑战消息
func GetChallenge(address string) (string, error) {
	if !common.IsHexAddress(address) {
		return "", fmt.Errorf("invalid Ethereum address")
	}

	// 生成 nonce
	nonce := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%d", time.Now().UnixNano()))))
	challenges[address] = nonce

	message := fmt.Sprintf(`web3blog.xyz wants you to sign in with your Ethereum account:
%s

This allows you to access your blog dashboard.
Nonce: %s`, address, nonce)

	return message, nil
}

// ValidateChallenge 验证签名是否匹配该地址的挑战消息
//
//	替代直接访问 challenges
func ValidateChallenge(address, signature, message string) bool {
	// 1. 检查是否有该地址的挑战
	expectedNonce, exists := challenges[address]
	if !exists {
		return false
	}

	// 2. 验证签名
	valid := validateSignature(address, signature, message)
	if !valid {
		return false
	}

	// 3. 可选：验证 nonce 是否在消息中（防伪造）
	if !strings.Contains(message, expectedNonce) {
		return false
	}

	return true
}

// ClearChallenge 登录成功后清除 nonce
func ClearChallenge(address string) {
	delete(challenges, address)
}

// 私有函数：验证签名
func validateSignature(address, signature, message string) bool {
	sig := common.FromHex(signature)
	if len(sig) == 65 {
		sig[64] -= 27
	}

	hash := crypto.Keccak256Hash([]byte(message))
	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return false
	}

	recovered := crypto.PubkeyToAddress(*pubKey)
	return strings.EqualFold(recovered.Hex(), address)
}
