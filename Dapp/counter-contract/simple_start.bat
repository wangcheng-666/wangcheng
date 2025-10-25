@echo off
chcp 65001 > nul

REM 设置简单模式下的颜色
color 0A

REM 创建必要的目录
mkdir contracts 2>nul
mkdir abi 2>nul
mkdir bin 2>nul
mkdir go-contract 2>nul

REM 创建简单的Counter合约
echo 创建智能合约...
echo pragma solidity ^0.8.0;^nevent CountIncreased(uint256 indexed newValue);^ncontract Counter {^n    uint256 public count = 0;^n    ^n    function increase() public {^n        count += 1;^n        emit CountIncreased(count);^n    }^n    ^n    function getCount() public view returns (uint256) {^n        return count;^n    }^n}> contracts\Counter.sol

REM 检查solcjs是否安装
where solcjs >nul 2>nul
if %errorlevel% neq 0 (
    echo 正在安装solcjs编译器...
    npm install -g solc
)

REM 编译合约
call :compile_contract

REM 检查abigen是否安装
where abigen >nul 2>nul
if %errorlevel% neq 0 (
    echo 正在安装abigen工具...
    go install github.com/ethereum/go-ethereum/cmd/abigen@latest
)

REM 生成Go绑定代码
call :generate_bindings

REM 创建go.mod文件
call :create_go_mod

REM 创建简单的Go交互程序
call :create_go_app

REM 完成提示
echo.
echo =================================================
echo 所有步骤已完成！
echo =================================================
echo 现在您可以：
echo 1. 编辑 main.go 文件，替换RPC URL和私钥
 echo 2. 运行: go mod tidy
 echo 3. 运行: go run main.go
 echo.
echo 注意：需要在Sepolia测试网络上有测试ETH
echo =================================================
pause
goto :eof

REM 编译合约函数
:compile_contract
echo 编译智能合约...
solcjs --abi --bin contracts\Counter.sol
move contracts_Counter_sol_Counter.abi abi\Counter.abi 2>nul
move contracts_Counter_sol_Counter.bin bin\Counter.bin 2>nul
echo 编译完成！ABI和字节码文件已生成到abi和bin目录
goto :eof

REM 生成Go绑定代码函数
:generate_bindings
echo 生成Go绑定代码...
abigen --abi abi\Counter.abi --bin bin\Counter.bin --pkg counter --out go-contract\counter.go
echo Go绑定代码已生成到go-contract\counter.go
goto :eof

REM 创建go.mod文件
:create_go_mod
if not exist go.mod (
    echo 创建go.mod文件...
echo module counter-contract^ngo 1.19^nrequire github.com/ethereum/go-ethereum v1.13.4^nreplace counter-contract/go-contract => ./go-contract> go.mod
)
goto :eof

REM 创建Go交互程序
:create_go_app
echo 创建Go交互程序...
echo package main^n^nimport (^n	"context"^n	"crypto/ecdsa"^n	"fmt"^n	"log"^n	"math/big"^n	"time"^n^n	"counter-contract/go-contract"^n^n	"github.com/ethereum/go-ethereum/accounts/abi/bind"^n	"github.com/ethereum/go-ethereum/common"^n	"github.com/ethereum/go-ethereum/crypto"^n	"github.com/ethereum/go-ethereum/ethclient"^n)^n^nfunc main() {^n	fmt.Println("=== 简单的以太坊合约交互程序 ===")^n^n	// 配置信息（需要修改）^n	rpcURL := "https://sepolia.infura.io/v3/YOUR_PROJECT_ID" // 替换为您的Infura项目ID^n	privateKeyHex := "YOUR_PRIVATE_KEY" // 替换为您的私钥^n^n	// 连接到Sepolia测试网^n	client, err := ethclient.Dial(rpcURL)^n	if err != nil {^n		log.Fatalf("无法连接到以太坊节点: %v", err)^n	}^n	defer client.Close()^n
echo 已连接到Sepolia测试网^n^n	// 获取链ID^n	chainID, err := client.ChainID(context.Background())^n	if err != nil {^n		log.Fatalf("获取链ID失败: %v", err)^n	}^n^n	// 解析私钥^n	privateKey, err := crypto.HexToECDSA(privateKeyHex)^n	if err != nil {^n		log.Fatalf("解析私钥失败: %v", err)^n	}^n^n	// 创建交易授权^n	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)^n	if err != nil {^n		log.Fatalf("创建交易授权失败: %v", err)^n	}^n	auth.GasLimit = uint64(300000)^n^n	// 部署合约^n	fmt.Println("开始部署合约...")^n	contractAddr, tx, instance, err := counter.DeployCounter(auth, client)^n	if err != nil {^n		log.Fatalf("部署合约失败: %v", err)^n	}^n^n	fmt.Printf("合约部署交易哈希: %s\n", tx.Hash().Hex())^n	fmt.Printf("合约地址: %s\n", contractAddr.Hex())^n	fmt.Println("等待交易确认...")^n	time.Sleep(30 * time.Second)^n^n	// 调用合约方法^n	fmt.Println("调用increase方法...")^n	tx, err = instance.Increase(auth)^n	if err != nil {^n		log.Fatalf("调用increase失败: %v", err)^n	}^n	fmt.Printf("交易哈希: %s\n", tx.Hash().Hex())^n	fmt.Println("等待交易确认...")^n	time.Sleep(30 * time.Second)^n^n	// 读取计数^n	count, err := instance.GetCount(nil)^n	if err != nil {^n		log.Fatalf("获取计数失败: %v", err)^n	}^n	fmt.Printf("当前计数: %d\n", count)^n^n	fmt.Println("\n程序执行完成！")^n}> main.go
goto :eof