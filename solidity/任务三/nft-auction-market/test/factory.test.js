// 工厂合约测试

import { expect } from 'chai';
import hre from 'hardhat';
const { ethers } = hre;
// 手动定义AddressZero常量，因为ethers.js v6中不再通过ethers.constants提供
const ADDRESS_ZERO = '0x0000000000000000000000000000000000000000';

describe('AuctionFactory', function () {
  let auctionFactory, marketNFT, deployer, user1, user2, feeRecipient;

  beforeEach(async function () {
    // 获取签名者
    [deployer, user1, user2, feeRecipient] = await ethers.getSigners();

    // 部署工厂合约（它会自动部署拍卖合约实现）
    const AuctionFactory = await ethers.getContractFactory('AuctionFactory');
    auctionFactory = await AuctionFactory.deploy();

    // 部署NFT合约
    const MarketNFT = await ethers.getContractFactory('MarketNFT');
    marketNFT = await MarketNFT.deploy();
  });

  describe('部署和初始化', function () {
    it('应该正确设置拍卖合约实现地址', async function () {
      const implementationAddress = await auctionFactory.auctionImplementation();
      expect(implementationAddress).to.not.equal(ADDRESS_ZERO);
    });

    it('应该设置正确的所有者', async function () {
      expect(await auctionFactory.owner()).to.equal(deployer.address);
    });

    it('应该从拍卖ID 0开始', async function () {
      expect(await auctionFactory.auctionCount()).to.equal(0);
    });
  });

  describe('创建拍卖合约', function () {
    it('应该允许所有者创建新的拍卖合约实例', async function () {
      const feePercent = 250; // 2.5%
      
      const tx = await auctionFactory.createAuctionContract(feePercent, feeRecipient.address);
      const receipt = await tx.wait();
      // 直接从合约状态中获取拍卖信息，避免处理事件
      // 获取创建的拍卖合约地址
      const auctionId = BigInt(await auctionFactory.auctionCount()) - BigInt(1);
      const auctionAddress = await auctionFactory.auctions(auctionId);
      
      expect(auctionId).to.equal(BigInt(0));
      expect(auctionAddress).to.not.equal(ADDRESS_ZERO);
      expect(await auctionFactory.auctionCount()).to.equal(1);
      expect(await auctionFactory.auctions(auctionId)).to.equal(auctionAddress);
    });

    it('应该允许创建多个拍卖合约实例', async function () {
      const feePercent1 = 250; // 2.5%
      const feePercent2 = 500; // 5%
      
      // 创建第一个拍卖合约
      await auctionFactory.createAuctionContract(feePercent1, feeRecipient.address);
      
      // 创建第二个拍卖合约
      const tx = await auctionFactory.createAuctionContract(feePercent2, feeRecipient.address);
      const receipt = await tx.wait();
      // 直接从合约状态中获取拍卖信息，避免处理事件
      // 获取创建的拍卖合约地址
      const auctionId = BigInt(await auctionFactory.auctionCount()) - BigInt(1);
      const auctionAddress = await auctionFactory.auctions(auctionId);
      
      expect(auctionId).to.equal(BigInt(1));
      expect(auctionAddress).to.not.equal(ADDRESS_ZERO);
      expect(await auctionFactory.auctionCount()).to.equal(2);
      expect(await auctionFactory.auctions(auctionId)).to.equal(auctionAddress);
    });

    it('不应该允许非所有者创建拍卖合约实例', async function () {
      const feePercent = 250; // 2.5%
      
      await expect(
        auctionFactory.connect(user1).createAuctionContract(feePercent, feeRecipient.address)
      ).to.be.revertedWith('Ownable: caller is not the owner');
    });

    it('应该正确初始化拍卖合约实例', async function () {
      const feePercent = 250; // 2.5%
      
      const tx = await auctionFactory.createAuctionContract(feePercent, feeRecipient.address);
      const receipt = await tx.wait();
      // 直接从合约状态中获取拍卖信息，避免处理事件
      // 获取创建的拍卖合约地址
      const auctionId = BigInt(await auctionFactory.auctionCount()) - BigInt(1);
      const auctionAddress = await auctionFactory.auctions(auctionId);
      
      // 获取拍卖合约实例
      const AuctionMarket = await ethers.getContractFactory('AuctionMarket');
      const auctionMarket = AuctionMarket.attach(auctionAddress);
      
      // 检查是否正确初始化
      expect(await auctionMarket.feePercent()).to.equal(feePercent);
      expect(await auctionMarket.feeRecipient()).to.equal(feeRecipient.address);
    });
  });

  describe('获取拍卖合约信息', function () {
    beforeEach(async function () {
      const feePercent = 250; // 2.5%
      
      await auctionFactory.createAuctionContract(feePercent, feeRecipient.address);
      await auctionFactory.createAuctionContract(feePercent, feeRecipient.address);
    });

    it('应该能够通过ID获取拍卖合约地址', async function () {
      const auctionAddress = await auctionFactory.auctions(0);
      expect(auctionAddress).to.not.equal(ADDRESS_ZERO);
    });

    it('应该能够获取所有拍卖合约地址', async function () {
      const allAuctions = await auctionFactory.getAllAuctionContracts();
      expect(allAuctions).to.have.lengthOf(2);
      expect(allAuctions[0]).to.not.equal(ADDRESS_ZERO);
      expect(allAuctions[1]).to.not.equal(ADDRESS_ZERO);
    });
  });

  describe('拍卖合约功能测试', function () {
    let auctionAddress, auctionMarket, tokenId;

    beforeEach(async function () {
      const feePercent = 250; // 2.5%
      
      // 创建拍卖合约实例
      const tx = await auctionFactory.createAuctionContract(feePercent, feeRecipient.address);
      const receipt = await tx.wait();
      // 直接从合约状态中获取拍卖信息，避免处理事件
      // 获取创建的拍卖合约地址
      const count = await auctionFactory.auctionCount();
      auctionAddress = await auctionFactory.auctions(BigInt(count) - BigInt(1));
      
      // 获取拍卖合约实例
      // 先检查auctionAddress是否有效
      expect(auctionAddress).to.not.equal(ADDRESS_ZERO);
      // 注意：在测试环境中，我们不需要再次初始化，因为createAuctionContract已经完成了初始化
      // 只需使用已创建的地址即可
      const AuctionMarket = await ethers.getContractFactory('AuctionMarket');
      auctionMarket = AuctionMarket.attach(auctionAddress);
      
      // 铸造NFT并转移给user1
      await marketNFT.mint(user1.address, 'https://example.com/nft/1');
      tokenId = 0;
      
      // user1授权拍卖合约转移NFT
      await marketNFT.connect(user1).approve(auctionMarket.address, tokenId);
    });

    it('应该能够使用工厂创建的拍卖合约进行拍卖', async function () {
      const startingPrice = ethers.utils.parseEther('1');
      const endTime = Math.floor(Date.now() / 1000) + 86400; // 24小时后
      const currency = ADDRESS_ZERO; // ETH
      const reservePrice = ethers.utils.parseEther('0.5');

      // user1创建拍卖
      const tx = await auctionMarket.connect(user1).createAuction(
        marketNFT.address,
        tokenId,
        startingPrice,
        endTime,
        currency,
        reservePrice
      );
      
      const receipt = await tx.wait();
      const event = receipt.events.find(event => event.event === 'AuctionCreated');
      const auctionId = event.args.auctionId;

      // user2出价
      const bidAmount = ethers.utils.parseEther('1.5');
      await auctionMarket.connect(user2).placeBid(auctionId, { value: bidAmount });

      // 检查出价是否成功
      const auction = await auctionMarket.getAuction(auctionId);
      expect(auction.highestBidder).to.equal(user2.address);
      expect(auction.highestBid).to.equal(bidAmount);
    });
  });
});