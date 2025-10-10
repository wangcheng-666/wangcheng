// scripts/verify-contracts.js
import { config } from 'dotenv';
import { execSync } from 'child_process';

// 加载.env文件
config();

// 设置环境变量
process.env.SEPOLIA_RPC_URL = process.env.SEPOLIA_RPC_URL || 'https://eth-sepolia.g.alchemy.com/v2/4HAF0v4P1UFGxMM4TR-we';
process.env.ETHERSCAN_API_KEY = process.env.ETHERSCAN_API_KEY || 'CQ8FYB5PKI963TK9U82Y84K3E4A6ZS4QYN';
// 设置SEPOLIA_PRIVATE_KEY变量（Hardhat验证时需要）
process.env.SEPOLIA_PRIVATE_KEY = process.env.PRIVATE_KEY || process.env.SEPOLIA_PRIVATE_KEY;

// 合约地址（使用之前部署的地址）
const nftContractAddress = '0x36CDcFC6Cd5867BA173905a4e54eF14e969184AE';
const auctionContractAddress = '0xe56822F4BDe1c513e05C3c242DB60dEEC26BDD44';
const factoryContractAddress = '0xde7611a561f5b53B98883bF14cf070266e9A725d';

console.log('开始验证合约...');
console.log('设置环境变量完成');

// 验证MarketNFT合约
try {
  console.log(`\n验证MarketNFT合约 (${nftContractAddress})...`);
  execSync(
    `npx hardhat verify --network sepolia ${nftContractAddress}`,
    { stdio: 'inherit' }
  );
  console.log('MarketNFT合约验证成功！');
} catch (error) {
  console.error('MarketNFT合约验证失败:', error.message);
}

// 验证AuctionMarket合约
try {
  console.log(`\n验证AuctionMarket合约 (${auctionContractAddress})...`);
  execSync(
    `npx hardhat verify --network sepolia ${auctionContractAddress}`,
    { stdio: 'inherit' }
  );
  console.log('AuctionMarket合约验证成功！');
} catch (error) {
  console.error('AuctionMarket合约验证失败:', error.message);
}

// 验证AuctionFactory合约
try {
  console.log(`\n验证AuctionFactory合约 (${factoryContractAddress})...`);
  execSync(
    `npx hardhat verify --network sepolia ${factoryContractAddress}`,
    { stdio: 'inherit' }
  );
  console.log('AuctionFactory合约验证成功！');
} catch (error) {
  console.error('AuctionFactory合约验证失败:', error.message);
}

console.log('\n合约验证过程完成！');
console.log('\n如果验证失败，请手动运行以下命令:');
console.log(`npx hardhat verify --network sepolia ${nftContractAddress}`);
console.log(`npx hardhat verify --network sepolia ${auctionContractAddress}`);
console.log(`npx hardhat verify --network sepolia ${factoryContractAddress}`);