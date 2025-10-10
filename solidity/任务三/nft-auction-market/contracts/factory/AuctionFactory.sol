// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import '../interfaces/IAuctionMarket.sol';
import '../auction/AuctionMarket.sol';
import '@openzeppelin/contracts/access/Ownable.sol';
import '@openzeppelin/contracts/proxy/Clones.sol';

/**
 * @title AuctionFactory
 * @dev 使用工厂模式创建和管理拍卖合约实例
 */
contract AuctionFactory is Ownable {
    // 拍卖合约实现的地址（用于克隆）
    address public immutable auctionImplementation;
    
    // 存储所有创建的拍卖合约
    mapping(uint256 => address) public auctions;
    uint256 public auctionCount;
    
    // 事件
    event AuctionContractCreated(
        uint256 indexed auctionId,
        address indexed auctionContract,
        address indexed creator
    );
    
    /**
     * @dev 构造函数，初始化拍卖合约实现地址
     */
    constructor() {
        // 部署拍卖合约实现
        auctionImplementation = address(new AuctionMarket());
    }
    
    /**
     * @dev 创建新的拍卖合约实例
     * @param feePercent 手续费比例（百分比，乘以100）
     * @param feeRecipient 手续费接收地址
     * @return auctionId 新创建的拍卖ID
     * @return auctionAddress 新创建的拍卖合约地址
     */
    function createAuctionContract(
        uint256 feePercent,
        address feeRecipient
    ) external onlyOwner returns (uint256, address) {
        // 使用克隆模式创建新的拍卖合约实例
        address payable auctionAddress = payable(Clones.clone(auctionImplementation));
        
        // 初始化拍卖合约
        AuctionMarket(auctionAddress).initialize(feePercent, feeRecipient);
        
        // 记录拍卖合约
        uint256 auctionId = auctionCount;
        auctions[auctionId] = auctionAddress;
        auctionCount++;
        
        emit AuctionContractCreated(auctionId, auctionAddress, msg.sender);
        
        return (auctionId, auctionAddress);
    }
    
    /**
     * @dev 获取拍卖合约地址
     * @param auctionId 拍卖ID
     * @return auctionAddress 拍卖合约地址
     */
    function getAuctionContract(uint256 auctionId) external view returns (address) {
        require(auctionId < auctionCount, 'Auction not found');
        return auctions[auctionId];
    }
    
    /**
     * @dev 获取所有拍卖合约地址
     * @return addresses 所有拍卖合约地址的数组
     */
    function getAllAuctionContracts() external view returns (address[] memory) {
        address[] memory addresses = new address[](auctionCount);
        for (uint256 i = 0; i < auctionCount; i++) {
            addresses[i] = auctions[i];
        }
        return addresses;
    }
}