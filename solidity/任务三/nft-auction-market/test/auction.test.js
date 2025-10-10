// 拍卖合约测试

import { expect } from 'chai';
import hre from 'hardhat';
const { ethers } = hre;
// 手动定义AddressZero常量，因为ethers.js v6中不再通过ethers.constants提供
const ADDRESS_ZERO = '0x0000000000000000000000000000000000000000';

describe('AuctionMarket', function () {
  let marketNFT, auctionMarket, deployer, seller, bidder1, bidder2, feeRecipient;
  let tokenId, auctionId;
  let marketNFTAddress, auctionMarketAddress;

  beforeEach(async function () {
    // 获取签名者
    [deployer, seller, bidder1, bidder2, feeRecipient] = await ethers.getSigners();

    // 部署NFT合约
    const MarketNFT = await ethers.getContractFactory('MarketNFT');
    marketNFT = await MarketNFT.deploy();
    marketNFTAddress = await marketNFT.getAddress();
    console.log('MarketNFT deployed at:', marketNFTAddress);

    // 部署拍卖合约
    const AuctionMarket = await ethers.getContractFactory('AuctionMarket');
    auctionMarket = await AuctionMarket.deploy();
    auctionMarketAddress = await auctionMarket.getAddress();
    console.log('AuctionMarket deployed at:', auctionMarketAddress);
    
    // 初始化拍卖合约（因为使用了UUPS代理模式）
    const feePercent = 250; // 2.5%
    await auctionMarket.initialize(feePercent, feeRecipient.address);
    console.log('AuctionMarket initialized with feePercent:', feePercent, 'and feeRecipient:', feeRecipient.address);

    // 检查初始化是否成功
    const initializedFeePercent = await auctionMarket.feePercent();
    const initializedFeeRecipient = await auctionMarket.feeRecipient();
    expect(initializedFeePercent).to.equal(feePercent);
    expect(initializedFeeRecipient).to.equal(feeRecipient.address);

    // 铸造NFT并转移给卖家
    await marketNFT.mint(seller.address, 'https://example.com/nft/1');
    tokenId = 0;
    console.log('NFT minted to seller:', seller.address, 'with tokenId:', tokenId);

    // 卖家授权拍卖合约转移NFT
    await marketNFT.connect(seller).approve(auctionMarketAddress, tokenId);
    console.log('NFT approved for transfer to AuctionMarket');
    
    // 验证授权
    const approved = await marketNFT.getApproved(tokenId);
    expect(approved).to.equal(auctionMarketAddress);
  });

  describe('基础功能测试', function () {
    it('应该能够创建拍卖', async function () {
      const startingPrice = ethers.parseEther('1');
      const endTime = Math.floor(Date.now() / 1000) + 86400; // 24小时后
      const currency = ADDRESS_ZERO; // ETH
      const reservePrice = ethers.parseEther('0.5');

      console.log('Creating auction with parameters:');
      console.log('- MarketNFT address:', marketNFTAddress);
      console.log('- Token ID:', tokenId);
      console.log('- Starting price:', ethers.formatEther(startingPrice), 'ETH');
      console.log('- Reserve price:', ethers.formatEther(reservePrice), 'ETH');
      
      // 创建拍卖
      const tx = await auctionMarket.connect(seller).createAuction(
        marketNFTAddress,
        tokenId,
        startingPrice,
        endTime,
        currency,
        reservePrice
      );
      
      console.log('Auction creation transaction sent:', tx.hash);
      
      // 等待交易确认
      const receipt = await tx.wait();
      console.log('Transaction confirmed in block:', receipt.blockNumber);
      
      // 简单方式获取拍卖ID（因为这是第一个拍卖）
      auctionId = 0;
      
      // 验证拍卖存在
      const auction = await auctionMarket.getAuction(auctionId);
      console.log('Auction retrieved:', auction);
      expect(auction.id).to.equal(auctionId);
    });
  });
});