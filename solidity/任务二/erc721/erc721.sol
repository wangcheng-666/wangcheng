// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyNFT is ERC721URIStorage, Ownable {
    uint256 private _nextTokenId;

    constructor() ERC721("MyAwesomeNFT", "MAN") Ownable(msg.sender) {
        _nextTokenId = 1;
    }

    //创建一个新的nft（metedata）授权给某个地址
    function mintNFT(address recipient, string memory tokenURI) 
        public 
        onlyOwner 
        returns (uint256)
    {
        uint256 tokenId = _nextTokenId++;
        _safeMint(recipient, tokenId); 
        _setTokenURI(tokenId, tokenURI);
        return tokenId;
    }
}