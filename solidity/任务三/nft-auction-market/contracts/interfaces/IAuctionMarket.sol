// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import '@openzeppelin/contracts/token/ERC721/IERC721.sol';
import '@openzeppelin/contracts/token/ERC20/IERC20.sol';

/**
 * @title IAuctionMarket
 * @dev 拍卖市场接口
 */
interface IAuctionMarket {
    struct Auction {        
        uint256 id;
        address nftContract;
        uint256 tokenId;
        address seller;
        uint256 startTime;
        uint256 endTime;
        uint256 startingPrice;
        uint256 reservePrice;
        address currency;
        address highestBidder;
        uint256 highestBid;
        bool isEnded;
    }

    // 事件
    event AuctionCreated(
        uint256 indexed auctionId,
        address indexed nftContract,
        uint256 indexed tokenId,
        address seller,
        uint256 startTime,
        uint256 endTime,
        uint256 startingPrice,
        address currency
    );

    event BidPlaced(
        uint256 indexed auctionId,
        address indexed bidder,
        uint256 amount
    );

    event AuctionEnded(
        uint256 indexed auctionId,
        address indexed winner,
        uint256 amount
    );

    // 函数
    function createAuction(
        address nftContract,
        uint256 tokenId,
        uint256 startingPrice,
        uint256 endTime,
        address currency,
        uint256 reservePrice
    ) external returns (uint256);

    function placeBid(uint256 auctionId) external payable;

    function endAuction(uint256 auctionId) external;

    function getAuction(uint256 auctionId) external view returns (Auction memory);

    function getPriceInUSD(address token, uint256 amount) external view returns (uint256);
}