# 区块链读写任务 - Sepolia测试网络交互

本项目实现了与以太坊Sepolia测试网络的基础交互功能，包括查询区块信息和发送以太币交易。

## 项目结构

```
blockchain-task/
├── cmd/
│   ├── block_query/         # 区块查询命令
│   └── send_transaction/    # 交易发送命令
├── internal/
│   ├── client/              # 以太坊客户端连接
│   ├── config/              # 配置管理
│   └── transaction/         # 交易处理
├── .env                     # 环境变量配置
├── .env.example             # 环境变量配置模板
├── .gitignore               # Git忽略文件
├── go.mod                   # Go模块定义
├── test_block_query.bat     # 区块查询测试脚本
└── test_send_transaction.bat # 交易发送测试脚本
```

## 环境搭建

### 1. 安装Go语言环境

请确保您已安装Go 1.16或更高版本。您可以从[Go官网](https://golang.org/dl/)下载并安装。

验证安装：
```bash
go version
```

### 2. 获取项目代码

```bash
# 克隆代码库（如果适用）
git clone <repository_url>
cd blockchain-task
```

### 3. 安装依赖

```bash
go mod tidy
```

## 什么是私钥和地址？

### 私钥（Private Key）

私钥是一串64位的十六进制字符串（不包含0x前缀），类似于：
`1a2b3c4d5e6f7g8h9i0j1k2l3m4n5o6p7q8r9s0t1u2v3w4x5y6z7a8b9c0d1e2f`

- 私钥是访问和控制您以太坊账户的唯一凭证
- 拥有私钥就意味着拥有对该账户中所有资产的控制权
- **请务必妥善保管您的私钥，不要与任何人分享！**

### 以太坊地址（Ethereum Address）

以太坊地址是一个42位的十六进制字符串（包含0x前缀），类似于：
`0x1a2b3c4d5e6f7g8h9i0j1k2l3m4n5o6p7q8r9s0t1u2v`

- 地址用于接收以太币和其他代币
- 您可以安全地与他人分享您的地址
- 地址是从私钥派生出来的，但无法从地址反推出私钥

## 如何获取私钥和地址

### 方法1：使用MetaMask钱包（推荐）

MetaMask是最常用的以太坊钱包之一，提供了简单的界面来管理您的账户。

1. **安装MetaMask浏览器扩展**
   - 访问 [MetaMask官网](https://metamask.io/) 下载并安装
   - 支持Chrome、Firefox、Edge等主流浏览器

2. **创建或导入账户**
   - 打开MetaMask并按照提示创建新钱包
   - 设置一个强密码并备份助记词（非常重要！）

3. **切换到Sepolia测试网络**
   - 点击MetaMask顶部的网络选择器（默认显示"Ethereum Mainnet"）
   - 选择"Sepolia Test Network"或点击"显示/隐藏测试网络"启用它

4. **导出私钥**
   - 点击MetaMask右上角的用户头像
   - 选择"账户详情"
   - 点击"导出私钥"
   - 输入您的MetaMask密码
   - 复制显示的私钥并妥善保管

### 方法2：使用命令行工具

如果您熟悉命令行，可以使用go-ethereum提供的geth工具生成账户：

```bash
# 安装geth工具
# 参考官方文档：https://geth.ethereum.org/docs/install-and-build/installing-geth

# 生成新账户
geth account new --keystore ./keystore

# 输入密码来保护您的账户
```

生成的私钥将保存在keystore目录中。

## 配置环境变量

获取私钥和地址后，您需要配置项目的.env文件：

1. 复制 `.env.example` 文件并重命名为 `.env`

2. 填写以下配置：
   ```bash
   # 复制配置文件
   copy .env.example .env
   # 编辑.env文件
   notepad .env
   ```

在`.env`文件中，您需要配置以下内容：

```env
# Infura API Key - 已配置
INFURA_API_KEY=0154bc1fc1804320969f1f66a2d27e41

# 以太坊私钥 - 请替换为您自己的私钥
# 格式：不包含0x前缀的64位十六进制字符串
PRIVATE_KEY=your_actual_private_key_here

# 接收地址 - 请替换为您要发送以太币的目标地址
# 格式：包含0x前缀的42位十六进制字符串
RECIPIENT_ADDRESS=your_recipient_address_here
```

## 获取Sepolia测试以太币

在发送交易前，您需要确保您的账户中有足够的Sepolia测试以太币：

1. 访问Sepolia水龙头网站：
   - [Infura Faucet](https://www.infura.io/faucet/sepolia)
   - [Alchemy Faucet](https://sepoliafaucet.com/)
   - [Chainstack Faucet](https://faucets.chainstack.com/sepolia-faucet/)

2. 粘贴您的Sepolia地址

3. 按照提示操作获取测试以太币

4. 等待几分钟，测试以太币将会出现在您的账户中

## 功能说明

### 1. 查询区块信息

该功能允许您查询Sepolia测试网络上的区块信息，包括区块哈希、时间戳、交易数量等。

#### 运行方式

使用提供的测试脚本：
```bash
test_block_query.bat
```

或者直接使用Go命令：
```bash
# 查询最新区块
go run cmd/block_query/main.go

# 查询指定区块
go run cmd/block_query/main.go 5000000
```

### 2. 发送交易

该功能允许您从配置的账户向指定接收地址发送Sepolia测试以太币。

#### 前提条件

1. 在`.env`文件中配置您的私钥和接收地址
2. 确保您的账户中有足够的Sepolia测试以太币
3. 了解私钥的重要性并妥善保管

#### 运行方式

使用提供的测试脚本：
```bash
test_send_transaction.bat
```

或者直接使用Go命令（默认转账1 ETH）：
```bash
go run cmd/send_transaction/main.go

# 指定转账金额（单位：wei）
go run cmd/send_transaction/main.go 500000000000000000  # 转账0.5 ETH
```

## 安全注意事项

- **永远不要在公共场合或代码库中泄露您的私钥**
- **不要将包含真实私钥的.env文件提交到GitHub等公共代码仓库**
- 使用强密码保护您的钱包和私钥存储
- 考虑使用硬件钱包（如Ledger、Trezor）来存储大额资产
- 测试网络上的私钥和地址可以随意生成，但主网上的私钥必须妥善保管

## 常见问题解答

### Q: 我可以使用主网的私钥和地址吗？
A: 可以，但在开发和测试阶段，强烈建议使用测试网络，以避免意外损失资金。

### Q: 私钥丢失了怎么办？
A: 如果您丢失了私钥且没有备份助记词，将无法恢复您的账户和资金。请务必妥善保管！

### Q: 如何检查我的账户余额？
A: 可以在MetaMask中查看，或使用Sepolia区块浏览器：https://sepolia.etherscan.io/

### Q: 交易需要多长时间才能确认？
A: 在Sepolia测试网络上，通常只需几秒钟到几分钟。确认时间取决于网络拥堵情况和您设置的燃气价格。

### Q: 为什么我的交易失败了？
A: 可能的原因包括：账户余额不足、燃气价格过低、nonce错误或网络问题。请检查错误消息以获取详细信息。

### Q: 如何查看交易状态？
A: 交易发送后，会输出交易哈希，您可以在Sepolia区块浏览器上查看交易状态：https://sepolia.etherscan.io/

## 资源链接

- [Sepolia区块浏览器](https://sepolia.etherscan.io/)
- [Go-ethereum文档](https://geth.ethereum.org/docs/)
- [以太坊开发者文档](https://ethereum.org/en/developers/)
- [MetaMask官网](https://metamask.io/)