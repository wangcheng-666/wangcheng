// scripts/find-and-deploy.js
import { ethers } from 'ethers';
import { config } from 'dotenv';
import { readFileSync, readdirSync, existsSync } from 'fs';
import { join } from 'path';

// 加载环境变量
config();

async function main() {
  console.log('开始部署NFT拍卖市场合约...');
  
  // 连接到Sepolia网络
  const provider = new ethers.JsonRpcProvider(process.env.SEPOLIA_RPC_URL);
  
  // 创建钱包实例
  const wallet = new ethers.Wallet(process.env.PRIVATE_KEY, provider);
  console.log('部署者地址:', wallet.address);
  
  // 检查部署者余额
  const balance = await provider.getBalance(wallet.address);
  console.log('部署者余额:', ethers.formatEther(balance), 'ETH');
  
  // 确保余额充足
  const minBalance = ethers.parseEther('0.01');
  if (balance < minBalance) {
    console.error('错误: 部署者余额不足，请获取更多Sepolia测试网ETH。');
    process.exit(1);
  }
  
  try {
    // 检查项目结构
    console.log('\n检查项目结构...');
    
    // 检查contracts目录
    const contractsDir = join(process.cwd(), 'contracts');
    if (existsSync(contractsDir)) {
      console.log('contracts目录存在，内容如下:');
      console.log(readdirSync(contractsDir, { withFileTypes: true })
        .map(dirent => dirent.isDirectory() ? `[目录] ${dirent.name}` : `[文件] ${dirent.name}`)
        .join('\n'));
    } else {
      console.error('错误: contracts目录不存在');
      process.exit(1);
    }
    
    // 检查artifacts目录
    const artifactsDir = join(process.cwd(), 'artifacts');
    if (existsSync(artifactsDir)) {
      console.log('\nartifacts目录存在，内容如下:');
      console.log(readdirSync(artifactsDir, { withFileTypes: true })
        .map(dirent => dirent.isDirectory() ? `[目录] ${dirent.name}` : `[文件] ${dirent.name}`)
        .join('\n'));
    } else {
      console.error('错误: artifacts目录不存在');
      process.exit(1);
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
    
    // 部署AuctionFactory合约
    console.log('\n部署工厂合约...');
    const factoryArtifact = JSON.parse(readFileSync(factoryArtifactPath, 'utf8'));
    const factoryFactory = new ethers.ContractFactory(factoryArtifact.abi, factoryArtifact.bytecode, wallet);
    const factoryContract = await factoryFactory.deploy();
    await factoryContract.waitForDeployment();
    const factoryAddress = await factoryContract.getAddress();
    console.log('工厂合约地址:', factoryAddress);
    
    console.log('\n合约部署完成！');
    console.log('\n提示: 合约部署完成后，可以使用以下命令验证合约:');
    console.log(`npx hardhat verify --network sepolia ${nftAddress}`);
    console.log(`npx hardhat verify --network sepolia ${auctionAddress}`);
    console.log(`npx hardhat verify --network sepolia ${factoryAddress}`);
  } catch (error) {
    console.error('部署过程中出错:', error);
    process.exit(1);
  }
}

main().catch((error) => {
  console.error(error);
  process.exit(1);
});