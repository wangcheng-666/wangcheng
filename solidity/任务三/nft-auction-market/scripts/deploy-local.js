// scripts/deploy-local.js
// 部署所有合约到本地Hardhat网络，方便在Remix中测试和交互

import { ethers } from 'ethers';
import { readFileSync, existsSync, writeFileSync } from 'fs';
import { join } from 'path';

// 添加延迟函数，确保交易有足够时间处理
const delay = (ms) => new Promise(resolve => setTimeout(resolve, ms));

async function main() {
  console.log('开始部署NFT拍卖市场合约到本地Hardhat网络...');
  
  // 连接到本地Hardhat节点
  const provider = new ethers.JsonRpcProvider('http://localhost:8545');
  
  // 获取第一个账户（Hardhat默认账户）
  const signers = await provider.listAccounts();
  console.log('可用账户:', signers);
  
  // 创建钱包实例
  // 注意：在实际使用中，您应该使用环境变量来存储私钥
  const wallet = new ethers.Wallet('0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80', provider);
  console.log('部署者地址:', wallet.address);
  
  // 检查部署者余额
  const balance = await provider.getBalance(wallet.address);
  console.log('部署者余额:', ethers.formatEther(balance), 'ETH');
  
  try {
    // 检查项目结构
    console.log('\n检查项目结构...');
    
    // 检查artifacts目录
    const artifactsDir = join(process.cwd(), 'artifacts');
    if (!existsSync(artifactsDir)) {
      console.log('artifacts目录不存在，尝试编译合约...');
      const { execSync } = await import('child_process');
      execSync('npx hardhat compile', { stdio: 'inherit' });
    }
    
    // 尝试找到正确的合约文件路径
    console.log('\n尝试查找合约文件...');
    
    // 搜索MarketNFT合约文件
    let nftArtifactPath = '';
    const searchPaths = [
      'contracts/MarketNFT.sol/MarketNFT.json',
      'contracts/nft/MarketNFT.sol/MarketNFT.json',
      'contracts/market/MarketNFT.sol/MarketNFT.json',
    ];
    
    for (const path of searchPaths) {
      const fullPath = join(artifactsDir, path);
      if (existsSync(fullPath)) {
        nftArtifactPath = fullPath;
        console.log('找到NFT合约文件:', nftArtifactPath);
        break;
      }
    }
    
    if (!nftArtifactPath) {
      console.error('错误: 未找到MarketNFT合约文件');
      process.exit(1);
    }
    
    // 部署MarketNFT合约
    console.log('\n部署NFT合约...');
    const nftArtifact = JSON.parse(readFileSync(nftArtifactPath, 'utf8'));
    const nftFactory = new ethers.ContractFactory(nftArtifact.abi, nftArtifact.bytecode, wallet);
    const nftContract = await nftFactory.deploy();
    await nftContract.waitForDeployment();
    const nftAddress = await nftContract.getAddress();
    console.log('NFT合约地址:', nftAddress);
    
    // 搜索AuctionMarket合约文件
    let auctionArtifactPath = '';
    const auctionSearchPaths = [
      'contracts/AuctionMarket.sol/AuctionMarket.json',
      'contracts/market/AuctionMarket.sol/AuctionMarket.json',
      'contracts/auction/AuctionMarket.sol/AuctionMarket.json',
    ];
    
    for (const path of auctionSearchPaths) {
      const fullPath = join(artifactsDir, path);
      if (existsSync(fullPath)) {
        auctionArtifactPath = fullPath;
        console.log('找到拍卖合约文件:', auctionArtifactPath);
        break;
      }
    }
    
    if (!auctionArtifactPath) {
      console.error('错误: 未找到AuctionMarket合约文件');
      process.exit(1);
    }
    
    // 部署AuctionMarket合约
    console.log('\n部署拍卖合约...');
    const auctionArtifact = JSON.parse(readFileSync(auctionArtifactPath, 'utf8'));
    const auctionFactory = new ethers.ContractFactory(auctionArtifact.abi, auctionArtifact.bytecode, wallet);
    const auctionContract = await auctionFactory.deploy();
    await auctionContract.waitForDeployment();
    const auctionAddress = await auctionContract.getAddress();
    
    // 初始化拍卖合约
    const feePercent = 250; // 2.5%
    await auctionContract.initialize(feePercent, wallet.address);
    console.log('拍卖合约地址:', auctionAddress);
    
    // 搜索AuctionFactory合约文件
    let factoryArtifactPath = '';
    const factorySearchPaths = [
      'contracts/AuctionFactory.sol/AuctionFactory.json',
      'contracts/factory/AuctionFactory.sol/AuctionFactory.json',
      'contracts/auction/AuctionFactory.sol/AuctionFactory.json',
    ];
    
    for (const path of factorySearchPaths) {
      const fullPath = join(artifactsDir, path);
      if (existsSync(fullPath)) {
        factoryArtifactPath = fullPath;
        console.log('找到工厂合约文件:', factoryArtifactPath);
        break;
      }
    }
    
    if (!factoryArtifactPath) {
      console.error('错误: 未找到AuctionFactory合约文件');
      process.exit(1);
    }
    
    // 增加延迟，避免nonce冲突
    console.log('\n等待5秒，确保之前的交易已确认...');
    await delay(5000);
    
    // 部署工厂合约
    console.log('\n部署工厂合约...');
    const factoryArtifact = JSON.parse(readFileSync(factoryArtifactPath, 'utf8'));
    
    // 获取当前nonce，确保使用正确的nonce
    const currentNonce = await provider.getTransactionCount(wallet.address, 'latest');
    console.log('当前nonce:', currentNonce);
    
    // 使用正确的nonce部署合约
    const factoryFactory = new ethers.ContractFactory(factoryArtifact.abi, factoryArtifact.bytecode, wallet);
    const factoryContract = await factoryFactory.deploy({ nonce: currentNonce });
    
    // 等待部署完成
    const deploymentReceipt = await factoryContract.waitForDeployment();
    const factoryAddress = await factoryContract.getAddress();
    console.log('工厂合约地址:', factoryAddress);
    console.log('工厂合约部署交易哈希:', deploymentReceipt.hash);
    
    console.log('\n合约部署完成！');
    console.log('\n以下是合约地址信息（请在Remix中使用这些地址）:');
    console.log(`MarketNFT: ${nftAddress}`);
    console.log(`AuctionMarket: ${auctionAddress}`);
    console.log(`AuctionFactory: ${factoryAddress}`);
    
    console.log('\n现在您可以：');
    console.log('1. 在Remix中连接到本地Hardhat网络 (http://localhost:8545)');
    console.log('2. 在Remix中加载合约ABI和地址');
    console.log('3. 与部署的合约进行交互和测试');
    
    console.log('\n提示: 合约部署信息已保存到deployed-contracts.txt文件中');
    
    // 保存部署信息到文件
    const deployInfo = `NFT拍卖市场合约部署信息 (本地Hardhat网络)\n` +
                      `部署时间: ${new Date().toLocaleString()}\n` +
                      `部署者地址: ${wallet.address}\n` +
                      `MarketNFT合约地址: ${nftAddress}\n` +
                      `AuctionMarket合约地址: ${auctionAddress}\n` +
                      `AuctionFactory合约地址: ${factoryAddress}\n`;
    
    const { writeFileSync } = await import('fs');
    writeFileSync('deployed-contracts.txt', deployInfo);
    
  } catch (error) {
    console.error('部署过程中出错:', error);
    process.exit(1);
  }
}

main().catch((error) => {
  console.error(error);
  process.exit(1);
});