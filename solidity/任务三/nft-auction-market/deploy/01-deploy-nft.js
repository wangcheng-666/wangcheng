// 部署NFT合约的脚本

import { ethers } from 'hardhat';

/**
 * @param {import('hardhat/types').HardhatRuntimeEnvironment} hre
 */
export default async function deployMarketNFT(hre) {
  const { getNamedAccounts, deployments, run } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  console.log('Deploying MarketNFT contract...');
  const marketNFT = await deploy('MarketNFT', {
    from: deployer,
    args: [],
    log: true,
    waitConfirmations: 1,
  });

  console.log('MarketNFT deployed to:', marketNFT.address);

  // 验证合约（如果支持）
  if (process.env.ETHERSCAN_API_KEY) {
    try {
      await run('verify:verify', {
        address: marketNFT.address,
        constructorArguments: [],
      });
    } catch (error) {
      console.log('Verification failed:', error.message);
    }
  }
}

deployMarketNFT.tags = ['MarketNFT'];