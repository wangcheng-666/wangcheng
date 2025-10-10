# NFT拍卖市场项目

一个基于Hardhat框架开发的NFT拍卖市场，支持Chainlink价格预言机、UUPS代理模式合约升级和工厂模式管理拍卖。

## 功能特点

- **NFT合约**：使用ERC721标准实现，支持铸造和转移
- **拍卖合约**：支持创建拍卖、出价（ETH/ERC20）、结束拍卖等功能
- **工厂模式**：使用工厂模式管理拍卖合约实例
- **Chainlink集成**：使用价格预言机计算资产美元价值
- **合约升级**：使用UUPS代理模式实现安全的合约升级

## 项目结构

```
├── .env                  # 环境变量配置
├── .env.example          # 环境变量配置示例
├── .gitignore            # Git忽略文件配置
├── contracts/            # 智能合约源码
│   ├── auction/          # 拍卖相关合约
│   ├── factory/          # 工厂相关合约
│   ├── interfaces/       # 接口定义
│   ├── libraries/        # 辅助库
│   └── nft/              # NFT相关合约
├── deploy/               # 部署脚本
├── scripts/              # 辅助脚本
├── test/                 # 测试文件
├── hardhat.config.js     # Hardhat配置文件
└── package.json          # 项目依赖
```

## 环境要求

- Node.js >= 16.0.0
- npm >= 7.0.0
- Visual Studio Code（推荐）
- MetaMask浏览器扩展

## 安装依赖

1. 打开终端，导航到项目目录：
   ```bash
   cd c:\Users\Administrator\Desktop\nft-auction-market
   ```

2. 安装项目依赖：
   ```bash
   npm install
   # 如果遇到依赖冲突，可以使用以下命令
   npm install --legacy-peer-deps
   # 有的包下载不下来
   npm install @nomicfoundation/hardhat-ethers@^3.0.0 --legacy-peer-deps

   npm install --save-dev @nomicfoundation/hardhat-network-helpers@^1.0.0 @nomicfoundation/hardhat-chai-matchers@^2.0.0 @nomicfoundation/hardhat-verify@^1.0.0 @typechain/ethers-v6@^0.4.0 @typechain/hardhat@^8.0.0 hardhat-gas-reporter@^1.0.8 solidity-coverage@^0.8.1 ts-node@>=8.0.0 typechain@^8.2.0 --legacy-peer-deps
   ```

3. 安装remixd工具（用于连接Remix IDE和本地文件系统）：
   ```bash
    #安装remixd插件
      npm install -g @remix-project/remixd
    #切换到项目文件夹
      cd c:\Users\Administrator\Desktop\nft-auction-market 
    #连接redix
      remixd -s . --remix-ide https://remix.ethereum.org
   ```

## 配置环境变量

1. 复制 `.env.example` 文件并命名为 `.env`：
   ```bash
   cp .env.example .env
   ```

2. 编辑 `.env` 文件，填入以下信息：
   ```
   # Sepolia测试网配置
   SEPOLIA_RPC_URL=https://eth-sepolia.g.alchemy.com/v2/[YOUR_ALCHEMY_API_KEY]
   PRIVATE_KEY=[YOUR_PRIVATE_KEY]
   ETHERSCAN_API_KEY=[YOUR_ETHERSCAN_API_KEY]
   ```

   - **SEPOLIA_RPC_URL**: 可以从Alchemy、Infura等服务获取
   - **PRIVATE_KEY**: 您的以太坊钱包私钥（用于部署合约）
   - **ETHERSCAN_API_KEY**: 从Etherscan网站获取

## 使用remixd连接Remix IDE

### 步骤1：启动remixd守护进程

remixd是Remix IDE的命令行工具，用于在本地文件系统和Remix IDE之间建立连接。

1. 打开一个新的终端窗口
2. 导航到项目目录：
   ```bash
   cd c:\Users\Administrator\Desktop\nft-auction-market
   ```
3. 启动remixd守护进程，连接当前目录：
   ```bash
   remixd -s . --remix-ide https://remix.ethereum.org
   ```
4. 成功启动后，您将看到类似以下输出：
   ```
   [WARN] You may now only use IDE at https://remix.ethereum.org to connect to this instance.
   [INFO] Shared folder: c:\Users\Administrator\Desktop\nft-auction-market
   [INFO] remixd is listening on 127.0.0.1:65520
   ```
5. **保持这个终端窗口打开**，不要关闭它！

### 步骤2：在Remix IDE中连接本地文件系统

1. 打开Remix IDE：https://remix.ethereum.org/
2. 在左侧边栏中，点击"Plugins"图标（通常是方块形状的图标）
3. 在插件列表中找到"Remixd"，点击旁边的"Activate"按钮激活它
4. 激活后，您会看到Remixd插件面板
5. 在面板中，点击"Connect to Localhost"按钮
6. 在弹出的确认对话框中，点击"Connect"按钮
7. 连接成功后，您可以在Remix IDE的文件浏览器中看到项目文件

### 步骤3：编译合约

1. 在Remix IDE左侧边栏中，点击"File Explorers"图标
2. 展开"localhost"文件夹，您应该能看到项目的所有文件
3. 导航到 `contracts/nft/MarketNFT.sol` 文件并点击打开
4. 在左侧边栏中，点击"Solidity Compiler"图标
5. 确保编译器版本设置为与合约兼容的版本（合约头部通常有版本声明，如 `pragma solidity ^0.8.28;`）
6. 点击"Compile MarketNFT.sol"按钮
7. 重复上述步骤编译其他合约文件：
   - `contracts/auction/AuctionMarket.sol`
   - `contracts/factory/AuctionFactory.sol`
   - `contracts/interfaces/IAuctionMarket.sol`
   - `contracts/libraries/PriceConverter.sol`

### 步骤4：启动本地Hardhat节点

为了在本地测试合约，我们需要启动一个Hardhat本地节点：

1. 打开另一个新的终端窗口
2. 导航到项目目录：
   ```bash
   cd c:\Users\Administrator\Desktop\nft-auction-market
   ```
3. 启动Hardhat本地节点：
   ```bash
   npx hardhat node
   ```
4. 成功启动后，您将看到10个测试账户和私钥的输出
5. **保持这个终端窗口打开**，不要关闭它！

### 步骤5：在Remix中部署合约到本地节点

1. 在Remix IDE左侧边栏中，点击"Deploy & Run Transactions"图标
2. 在"Environment"下拉菜单中，选择"Hardhat Provider"
3. 确保"Account"下拉菜单中显示的是Hardhat提供的测试账户之一
4. 从"Contract"下拉菜单中选择您要部署的合约（例如：`MarketNFT`）
5. 点击"Deploy"按钮
6. 部署成功后，您将在"Deployed Contracts"部分看到您的合约实例
7. 重复上述步骤部署其他合约

## 部署到Sepolia测试网

### 步骤1：准备工作

1. 确保您的`.env`文件已正确配置
2. 获取一些Sepolia测试网ETH（可以从水龙头获取：https://sepoliafaucet.com/）
3. 确保您的MetaMask钱包已配置Sepolia网络并拥有测试ETH

### 步骤2：使用脚本部署到Sepolia

1. 打开终端窗口
2. 导航到项目目录：
   ```bash
   cd c:\Users\Administrator\Desktop\nft-auction-market
   ```
3. 运行部署脚本：
   ```bash
   node scripts/find-and-deploy.js
   ```
4. 部署过程中，您将看到合约部署的进度和结果
5. 部署成功后，您会看到合约地址输出到控制台

### 步骤3：在Remix中连接Sepolia网络

1. 在Remix IDE的"Deploy & Run Transactions"面板中
2. 从"Environment"下拉菜单中选择"Injected Provider - MetaMask"
3. MetaMask会弹出确认对话框，点击"Next"和"Connect"按钮
4. 确保MetaMask已切换到Sepolia网络
5. 现在您可以在Remix中与已部署的Sepolia合约交互

### 步骤4：验证合约（可选）

1. 在部署完成后，运行验证脚本：
   ```bash
   node scripts/verify-contracts.js
   ```
2. 验证成功后，您可以在Etherscan上查看已验证的合约代码

## 在Remix中与合约交互

### MarketNFT合约交互

1. 确保您已部署了MarketNFT合约
2. 在"Deployed Contracts"部分，展开MarketNFT合约
3. 您可以调用以下主要函数：
   - `mint(address to, string memory tokenURI)`: 铸造新的NFT
   - `mintBatch(address to, string[] memory tokenURIs)`: 批量铸造NFT
   - `tokenURI(uint256 tokenId)`: 获取NFT的URI

### AuctionMarket合约交互

1. 确保您已部署了AuctionMarket合约
2. 在"Deployed Contracts"部分，展开AuctionMarket合约
3. 您可以调用以下主要函数：
   - `createAuction(address nftContract, uint256 tokenId, uint256 startingPrice, uint256 endTime, address currency, uint256 reservePrice)`: 创建新的拍卖
   - `placeBid(uint256 auctionId)`: 参与拍卖出价
   - `endAuction(uint256 auctionId)`: 结束拍卖并处理结果

### AuctionFactory合约交互

1. 确保您已部署了AuctionFactory合约
2. 在"Deployed Contracts"部分，展开AuctionFactory合约
3. 您可以调用以下主要函数：
   - `createAuctionContract(uint256 feePercent, address feeRecipient)`: 创建新的拍卖合约
   - `getAllAuctionContracts()`: 获取所有拍卖合约地址

## 常见问题和故障排除

### 1. remixd连接问题

**症状**：无法连接到remixd守护进程，提示"Cannot connect to the remixd daemon."

**解决方案**：
- 确保remixd守护进程正在运行（检查启动remixd的终端窗口）
- 确认您使用的是正确的命令：`remixd -s . --remix-ide https://remix.ethereum.org`
- 检查防火墙设置，确保65520端口未被阻止
- 尝试关闭所有浏览器窗口，然后重新打开Remix IDE
- 尝试使用管理员权限运行终端

### 2. Hardhat节点连接问题

**症状**：无法连接到localhost或Hardhat节点

**解决方案**：
- 确保Hardhat节点正在运行（检查启动节点的终端窗口）
- 确认节点运行在正确的端口上（默认是8545）
- 检查防火墙设置，确保8545端口未被阻止

### 3. 合约部署失败

**症状**：部署合约时出现错误

**解决方案**：
- 检查账户余额是否充足（特别是在Sepolia测试网）
- 确认.env文件配置正确
- 查看终端输出的错误信息，根据具体错误进行排查

### 4. Remix中看不到本地文件

**症状**：连接remixd后，在Remix IDE中看不到项目文件

**解决方案**：
- 确认remixd命令中的路径正确（使用`.`表示当前目录）
- 检查项目目录权限，确保remixd有足够的访问权限
- 尝试重新启动remixd守护进程

## 安全注意事项

- 在生产环境中，请确保正确配置私钥和API密钥
- 不要在公共场合或不安全的地方分享包含私钥的信息
- 建议在部署前进行全面的安全审计
- 使用remixd时，请注意您正在共享本地文件系统，确保只共享必要的项目目录

## 许可证

[MIT License]