// NFT合约测试

import { expect } from 'chai';
import hre from 'hardhat';
const { ethers } = hre;

describe('MarketNFT', function () {
  let marketNFT, deployer, user1, user2;

  beforeEach(async function () {
    // 获取签名者
    [deployer, user1, user2] = await ethers.getSigners();

    // 部署NFT合约
    const MarketNFT = await ethers.getContractFactory('MarketNFT');
    marketNFT = await MarketNFT.deploy();
  });

  describe('部署和初始化', function () {
    it('应该正确设置名称和符号', async function () {
      expect(await marketNFT.name()).to.equal('MarketNFT');
      expect(await marketNFT.symbol()).to.equal('MNFT');
    });

    it('应该设置正确的所有者', async function () {
      expect(await marketNFT.owner()).to.equal(deployer.address);
    });

    it('应该从tokenId 0开始', async function () {
      expect(await marketNFT.getNextTokenId()).to.equal(0);
    });
  });

  describe('铸造功能', function () {
    it('应该允许所有者铸造新NFT', async function () {
      const tokenURI = 'https://example.com/nft/1';
      const tx = await marketNFT.mint(user1.address, tokenURI);
      await tx.wait();

      expect(await marketNFT.ownerOf(0)).to.equal(user1.address);
      expect(await marketNFT.tokenURI(0)).to.equal(tokenURI);
      expect(await marketNFT.getNextTokenId()).to.equal(1);
    });

    it('应该允许批量铸造NFT', async function () {
      const tokenURIs = ['https://example.com/nft/1', 'https://example.com/nft/2', 'https://example.com/nft/3'];
      const tx = await marketNFT.mintBatch(user1.address, tokenURIs);
      await tx.wait();

      expect(await marketNFT.ownerOf(0)).to.equal(user1.address);
      expect(await marketNFT.ownerOf(1)).to.equal(user1.address);
      expect(await marketNFT.ownerOf(2)).to.equal(user1.address);
      expect(await marketNFT.tokenURI(0)).to.equal(tokenURIs[0]);
      expect(await marketNFT.tokenURI(1)).to.equal(tokenURIs[1]);
      expect(await marketNFT.tokenURI(2)).to.equal(tokenURIs[2]);
      expect(await marketNFT.getNextTokenId()).to.equal(3);
    });

    it('不应该允许非所有者铸造NFT', async function () {
      const tokenURI = 'https://example.com/nft/1';
      await expect(marketNFT.connect(user1).mint(user1.address, tokenURI)).to.be.revertedWith('Ownable: caller is not the owner');
    });

    it('应该允许所有者设置Token URI', async function () {
      const tokenURI1 = 'https://example.com/nft/1';
      const tokenURI2 = 'https://example.com/nft/updated';
      
      await marketNFT.mint(user1.address, tokenURI1);
      await marketNFT.setTokenURI(0, tokenURI2);
      
      expect(await marketNFT.tokenURI(0)).to.equal(tokenURI2);
    });

    it('不应该允许设置不存在的Token的URI', async function () {
      const tokenURI = 'https://example.com/nft/1';
      await expect(marketNFT.setTokenURI(999, tokenURI)).to.be.revertedWith('ERC721Metadata: URI set for nonexistent token');
    });
  });

  describe('NFT转移功能', function () {
    beforeEach(async function () {
      const tokenURI = 'https://example.com/nft/1';
      await marketNFT.mint(user1.address, tokenURI);
    });

    it('应该允许所有者转移NFT', async function () {
      await marketNFT.connect(user1).transferFrom(user1.address, user2.address, 0);
      expect(await marketNFT.ownerOf(0)).to.equal(user2.address);
    });

    it('应该允许授权地址转移NFT', async function () {
      await marketNFT.connect(user1).approve(user2.address, 0);
      await marketNFT.connect(user2).transferFrom(user1.address, user2.address, 0);
      expect(await marketNFT.ownerOf(0)).to.equal(user2.address);
    });
  });
});