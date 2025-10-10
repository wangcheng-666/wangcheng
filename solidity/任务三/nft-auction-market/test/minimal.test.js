import { expect } from 'chai';
import hre from 'hardhat';
const { ethers } = hre;

// 手动定义AddressZero常量，因为ethers.js v6中不再通过ethers.constants提供
const ADDRESS_ZERO = '0x0000000000000000000000000000000000000000';

describe('Minimal AuctionMarket Test', function () {
  let auctionMarket, deployer, feeRecipient;

  beforeEach(async function () {
    // 获取签名者
    const signers = await ethers.getSigners();
    console.log('Signers count:', signers.length);
    
    deployer = signers[0];
    feeRecipient = signers[1];
    
    console.log('Deployer address:', deployer.address);
    console.log('Fee recipient address:', feeRecipient.address);

    // 部署拍卖合约
    console.log('Deploying AuctionMarket...');
    const AuctionMarket = await ethers.getContractFactory('AuctionMarket');
    auctionMarket = await AuctionMarket.deploy();
    console.log('AuctionMarket deployed to:', auctionMarket.address);
    
    // 初始化拍卖合约（因为使用了UUPS代理模式）
    const feePercent = 250; // 2.5%
    console.log('Initializing AuctionMarket with feePercent:', feePercent, 'and feeRecipient:', feeRecipient.address);
    await auctionMarket.initialize(feePercent, feeRecipient.address);
    console.log('AuctionMarket initialized successfully');
  });

  it('should be properly initialized', async function () {
    const feePercent = await auctionMarket.feePercent();
    const recipient = await auctionMarket.feeRecipient();
    
    console.log('Stored feePercent:', feePercent.toString());
    console.log('Stored feeRecipient:', recipient);
    
    expect(feePercent).to.equal(250);
    expect(recipient).to.equal(feeRecipient.address);
  });
});