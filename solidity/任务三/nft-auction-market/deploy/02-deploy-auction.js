// 部署拍卖合约的脚本

import { ethers } from 'hardhat';

/**
 * @param {import('hardhat/types').HardhatRuntimeEnvironment} hre
 */
export default async function deployAuctionMarket(hre) {
  const { getNamedAccounts, deployments, run } = hre;
  const { deploy } = deployments;
  const { deployer, feeRecipient } = await getNamedAccounts();

  // 手续费比例：250 表示 2.5%（因为存储的是乘以100的百分比）
  const feePercent = 250;
  // 手续费接收地址，如果未配置则使用部署者地址
  const recipient = feeRecipient || deployer;

  console.log('Deploying AuctionMarket contract...');
  console.log('Fee percent:', feePercent, '(', feePercent / 100, '%)');
  console.log('Fee recipient:', recipient);

  const auctionMarket = await deploy('AuctionMarket', {
    from: deployer,
    args: [], // 使用initialize代替构造函数参数
    log: true,
    waitConfirmations: 1,
    proxy: {
      proxyContract: 'UUPS',
      execute: {
        init: {
          methodName: 'initialize',
          args: [feePercent, recipient],
        },
      },
    },
  });

  console.log('AuctionMarket deployed to:', auctionMarket.address);
  console.log('AuctionMarket implementation deployed to:', auctionMarket.implementation);

  // 验证合约（如果支持）
  if (process.env.ETHERSCAN_API_KEY) {
    try {
      await run('verify:verify', {
        address: auctionMarket.implementation,
        constructorArguments: [],
      });
    } catch (error) {
      console.log('Verification failed:', error.message);
    }
  }
}

deployAuctionMarket.tags = ['AuctionMarket'];
deployAuctionMarket.dependencies = [];