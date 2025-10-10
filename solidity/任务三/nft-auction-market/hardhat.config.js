import '@nomicfoundation/hardhat-ethers';
import '@nomicfoundation/hardhat-toolbox';
import 'hardhat-deploy';

/** @type import('hardhat/config').HardhatUserConfig */
export default {
  solidity: {
    version: '0.8.20',
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
    },
  },
  networks: {
    hardhat: {
      chainId: 1337,
      // 允许外部连接到本地节点
      allowUnlimitedContractSize: true,
      gas: 12000000,
      blockGasLimit: 0x1fffffffffffff,
      // 配置初始账户余额，方便测试
      accounts: {
        accountsBalance: '10000000000000000000000'
      }
    },
    localhost: {
      url: 'http://localhost:8545',
      chainId: 1337,
      // 可以使用与hardhat网络相同的账户
      accounts: ['85ace932e850141a2160f10d063b70ad8b094739d7658f07229144defc4108df']
    },
    sepolia: {
      url: 'https://eth-sepolia.g.alchemy.com/v2/4HAF0v4P1UFGxMM4TR-we',
      accounts: ['85ace932e850141a2160f10d063b70ad8b094739d7658f07229144defc4108df'],
    },
  },
  etherscan: {
    apiKey: 'CQ8FYB5PKI963TK9U82Y84K3E4A6ZS4QYN',
  },
  paths: {
    sources: './contracts',
    tests: './test',
    cache: './cache',
    artifacts: './artifacts',
    deploy: './deploy',
  },
  mocha: {
    timeout: 40000,
  },
};