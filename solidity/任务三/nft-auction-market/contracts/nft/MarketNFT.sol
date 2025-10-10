// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import '@openzeppelin/contracts/token/ERC721/ERC721.sol';
import '@openzeppelin/contracts/access/Ownable.sol';

/**
 * @title MarketNFT
 * @dev 拍卖市场使用的NFT合约
 */
contract MarketNFT is ERC721, Ownable {
    uint256 private _tokenIdCounter;
    mapping(uint256 => string) private _tokenURIs;

    constructor() ERC721('MarketNFT', 'MNFT') {
        _tokenIdCounter = 0;
    }

    /**
     * @dev 铸造新的NFT
     * @param to 接收NFT的地址
     * @param tokenURI NFT的URI
     * @return tokenId 新铸造的NFT ID
     */
    function mint(address to, string memory tokenURI) public onlyOwner returns (uint256) {
        uint256 tokenId = _tokenIdCounter;
        _mint(to, tokenId);
        _tokenURIs[tokenId] = tokenURI;
        _tokenIdCounter++;
        return tokenId;
    }

    /**
     * @dev 批量铸造NFT
     * @param to 接收NFT的地址
     * @param tokenURIs NFT的URI数组
     * @return tokenIds 新铸造的NFT ID数组
     */
    function mintBatch(address to, string[] memory tokenURIs) public onlyOwner returns (uint256[] memory) {
        uint256[] memory tokenIds = new uint256[](tokenURIs.length);
        for (uint256 i = 0; i < tokenURIs.length; i++) {
            tokenIds[i] = mint(to, tokenURIs[i]);
        }
        return tokenIds;
    }

    /**
     * @dev 获取NFT的URI
     * @param tokenId NFT的ID
     * @return tokenURI NFT的URI
     */
    function tokenURI(uint256 tokenId) public view override returns (string memory) {
        require(_exists(tokenId), 'ERC721Metadata: URI query for nonexistent token');
        return _tokenURIs[tokenId];
    }

    /**
     * @dev 设置NFT的URI
     * @param tokenId NFT的ID
     * @param tokenURI NFT的新URI
     */
    function setTokenURI(uint256 tokenId, string memory tokenURI) public onlyOwner {
        require(_exists(tokenId), 'ERC721Metadata: URI set for nonexistent token');
        _tokenURIs[tokenId] = tokenURI;
    }

    /**
     * @dev 获取当前可用的下一个Token ID
     * @return 下一个Token ID
     */
    function getNextTokenId() public view returns (uint256) {
        return _tokenIdCounter;
    }
}