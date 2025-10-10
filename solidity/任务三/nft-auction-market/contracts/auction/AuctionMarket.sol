// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import '../interfaces/IAuctionMarket.sol';
import '../libraries/PriceConverter.sol';
import '@openzeppelin/contracts/token/ERC721/IERC721.sol';
import '@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol';
import '@openzeppelin/contracts/token/ERC20/IERC20.sol';
import '@openzeppelin/contracts/access/Ownable.sol';
import '@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol';
import '@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol';

/**
 * @title AuctionMarket
 * @dev NFT拍卖市场的主要合约实现，支持NFT拍卖、出价和结束拍卖等功能
 */
contract AuctionMarket is IAuctionMarket, IERC721Receiver, Initializable, UUPSUpgradeable, OwnableUpgradeable {
    using PriceConverter for uint256;

    // 存储拍卖信息
    mapping(uint256 => Auction) private _auctions;
    uint256 private _auctionCounter;
    
    // 手续费比例（百分比，乘以100）
    uint256 public feePercent;
    // 手续费接收地址
    address public feeRecipient;
    
    // 实现UUPS代理模式所需的授权升级函数
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    /**
     * @dev 初始化函数，替代构造函数
     */
    function initialize(uint256 _feePercent, address _feeRecipient) public initializer {
        __Ownable_init();
        __UUPSUpgradeable_init();
        _auctionCounter = 0;
        feePercent = _feePercent;
        feeRecipient = _feeRecipient;
    }

    /**
     * @dev 创建NFT拍卖
     * @param nftContract NFT合约地址
     * @param tokenId NFT的Token ID
     * @param startingPrice 起拍价格
     * @param endTime 拍卖结束时间戳
     * @param currency 出价使用的货币地址（0地址表示使用ETH）
     * @param reservePrice 保留价格（可选，低于此价格不会成交）
     * @return auctionId 新创建的拍卖ID
     */
    function createAuction(
        address nftContract,
        uint256 tokenId,
        uint256 startingPrice,
        uint256 endTime,
        address currency,
        uint256 reservePrice
    ) external returns (uint256) {
        require(nftContract != address(0), 'NFT contract address cannot be zero');
        require(startingPrice > 0, 'Starting price must be greater than zero');
        require(endTime > block.timestamp, 'End time must be in the future');

        IERC721 nft = IERC721(nftContract);
        require(nft.ownerOf(tokenId) == msg.sender, 'You are not the owner of the NFT');
        require(nft.getApproved(tokenId) == address(this) || nft.isApprovedForAll(msg.sender, address(this)), 'Contract not approved to transfer NFT');

        uint256 auctionId = _auctionCounter;
        _auctions[auctionId] = Auction({
            id: auctionId,
            nftContract: nftContract,
            tokenId: tokenId,
            seller: msg.sender,
            startTime: block.timestamp,
            endTime: endTime,
            startingPrice: startingPrice,
            reservePrice: reservePrice,
            currency: currency,
            highestBidder: address(0),
            highestBid: 0,
            isEnded: false
        });

        _auctionCounter++;

        // 将NFT转移到合约中
        nft.safeTransferFrom(msg.sender, address(this), tokenId);

        emit AuctionCreated(
            auctionId,
            nftContract,
            tokenId,
            msg.sender,
            block.timestamp,
            endTime,
            startingPrice,
            currency
        );

        return auctionId;
    }

    /**
     * @dev 出价参与拍卖
     * @param auctionId 拍卖ID
     */
    function placeBid(uint256 auctionId) external payable {
        Auction storage auction = _auctions[auctionId];
        require(!auction.isEnded, 'Auction has ended');
        require(block.timestamp < auction.endTime, 'Auction has ended');
        require(auction.seller != msg.sender, 'Seller cannot bid');

        uint256 bidAmount;
        if (auction.currency == address(0)) {
            // 使用ETH出价
            bidAmount = msg.value;
        } else {
            // 使用ERC20代币出价
            // 注意：需要提前授权合约转移代币
            require(msg.value == 0, 'Cannot send ETH when bidding with ERC20');
            // 这里简化处理，实际应用中应该有前端传入出价金额
            revert('ERC20 bidding is not implemented in this simplified version');
        }

        require(bidAmount > auction.highestBid, 'Bid amount must be higher than current highest bid');
        require(bidAmount >= auction.startingPrice, 'Bid amount must be at least starting price');

        // 退还之前最高出价者的资金
        if (auction.highestBidder != address(0)) {
            if (auction.currency == address(0)) {
                // 退还ETH
                (bool success, ) = auction.highestBidder.call{value: auction.highestBid}('');
                require(success, 'Refund failed');
            } else {
                // 退还ERC20代币
                revert('ERC20 refund is not implemented in this simplified version');
            }
        }

        // 更新最高出价
        auction.highestBidder = msg.sender;
        auction.highestBid = bidAmount;

        emit BidPlaced(auctionId, msg.sender, bidAmount);
    }

    /**
     * @dev 结束拍卖
     * @param auctionId 拍卖ID
     */
    function endAuction(uint256 auctionId) external {
        Auction storage auction = _auctions[auctionId];
        require(!auction.isEnded, 'Auction already ended');
        require(block.timestamp >= auction.endTime, 'Auction has not ended yet');

        auction.isEnded = true;

        IERC721 nft = IERC721(auction.nftContract);

        if (auction.highestBidder != address(0) && auction.highestBid >= auction.reservePrice) {
            // 拍卖成功，转移NFT给最高出价者
            nft.safeTransferFrom(address(this), auction.highestBidder, auction.tokenId);

            // 计算手续费
            uint256 fee = auction.highestBid * feePercent / 10000; // feePercent是乘以100的百分比，所以除以10000
            uint256 sellerAmount = auction.highestBid - fee;

            // 支付卖家和手续费
            if (auction.currency == address(0)) {
                // 使用ETH支付
                (bool sellerSuccess, ) = auction.seller.call{value: sellerAmount}('');
                require(sellerSuccess, 'Payment to seller failed');
                
                (bool feeSuccess, ) = feeRecipient.call{value: fee}('');
                require(feeSuccess, 'Payment to fee recipient failed');
            } else {
                // 使用ERC20代币支付
                revert('ERC20 payment is not implemented in this simplified version');
            }

            emit AuctionEnded(auctionId, auction.highestBidder, auction.highestBid);
        } else {
            // 拍卖失败，退还NFT给卖家
            nft.safeTransferFrom(address(this), auction.seller, auction.tokenId);
            emit AuctionEnded(auctionId, address(0), 0);
        }
    }

    /**
     * @dev 获取拍卖信息
     * @param auctionId 拍卖ID
     * @return auction 拍卖信息
     */
    function getAuction(uint256 auctionId) external view returns (Auction memory) {
        return _auctions[auctionId];
    }

    /**
     * @dev 获取拍卖数量
     * @return 拍卖总数
     */
    function getAuctionCount() external view returns (uint256) {
        return _auctionCounter;
    }

    /**
     * @dev 更新手续费比例
     * @param _feePercent 新的手续费比例
     */
    function updateFeePercent(uint256 _feePercent) external onlyOwner {
        require(_feePercent <= 10000, 'Fee percent cannot exceed 100%');
        feePercent = _feePercent;
    }

    /**
     * @dev 更新手续费接收地址
     * @param _feeRecipient 新的手续费接收地址
     */
    function updateFeeRecipient(address _feeRecipient) external onlyOwner {
        require(_feeRecipient != address(0), 'Fee recipient cannot be zero address');
        feeRecipient = _feeRecipient;
    }

    /**
     * @dev 获取价格（以美元计）
     * @param token 代币地址（0地址表示ETH）
     * @param amount 代币数量
     * @return 美元价格（乘以10^8）
     */
    function getPriceInUSD(address token, uint256 amount) external view returns (uint256) {
        if (token == address(0)) {
            // ETH价格转换
            return amount.convertEthToUSD();
        } else {
            // ERC20价格转换（简化处理，实际应用中需要对应代币的价格预言机）
            revert('ERC20 price conversion is not implemented in this simplified version');
        }
    }

    /**
     * @dev 实现IERC721Receiver接口的onERC721Received函数
     * 允许合约接收NFT
     */
    function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes calldata data
    ) external pure override returns (bytes4) {
        // 返回IERC721Receiver.onERC721Received.selector，表明合约接受NFT
        return IERC721Receiver.onERC721Received.selector;
    }

    // 防止合约被意外锁定资金
    receive() external payable {}
}