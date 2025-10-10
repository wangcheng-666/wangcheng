// 部署工厂合约的脚本

import { ethers } from 'hardhat';

/**
 * @param {import('hardhat/types').HardhatRuntimeEnvironment} hre
 */
export default async function deployAuctionFactory(hre) {
  const { getNamedAccounts, deployments, run } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  console.log('Deploying AuctionFactory contract...');
  const auctionFactory = await deploy('AuctionFactory', {
    from: deployer,
    args: [],
    log: true,
    waitConfirmations: 1,
  });

  console.log('AuctionFactory deployed to:', auctionFactory.address);

  // 验证合约（如果支持）
  if (process.env.ETHERSCAN_API_KEY) {
    try {
      await run('verify:verify', {
        address: auctionFactory.address,
        constructorArguments: [],
      });
    } catch (error) {
      console.log('Verification failed:', error.message);
    }
  }
}

deployAuctionFactory.tags = ['AuctionFactory'];
deployAuctionFactory.dependencies = ['AuctionMarket'];