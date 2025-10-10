// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import '@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol';
import '@openzeppelin/contracts/utils/math/SafeMath.sol';

/**
 * @title PriceConverter
 * @dev 价格转换库，用于获取ERC20和以太坊到美元的价格
 */
library PriceConverter {
    using SafeMath for uint256;

    // 链上的ETH/USD价格预言机地址（Goerli测试网）
    address public constant ETH_USD_PRICE_FEED = 0xD4a33860578De61DBAbDc8BFdb98FD742fA7028e;

    /**
     * @dev 获取最新的ETH价格（以美元计价）
     * @return ethPriceInUSD ETH的最新价格（单位：美元，乘以10^8）
     */
    function getLatestEthPrice() internal view returns (uint256) {
        AggregatorV3Interface priceFeed = AggregatorV3Interface(ETH_USD_PRICE_FEED);
        (, int256 price, , , ) = priceFeed.latestRoundData();
        // 价格已经乘以10^8，转换为uint256
        return uint256(price);
    }

    /**
     * @dev 将ETH金额转换为美元
     * @param ethAmount ETH金额（单位：wei）
     * @return usdAmount 对应的美元金额（单位：美元，乘以10^8）
     */
    function convertEthToUSD(uint256 ethAmount) internal view returns (uint256) {
        uint256 ethPrice = getLatestEthPrice();
        // ethAmount (wei) * ethPrice (usd * 10^8) / 1e18 (wei/eth) = usd * 10^8
        return ethAmount.mul(ethPrice).div(1e18);
    }

    /**
     * @dev 将美元金额转换为ETH
     * @param usdAmount 美元金额（单位：美元，乘以10^8）
     * @return ethAmount 对应的ETH金额（单位：wei）
     */
    function convertUSDToEth(uint256 usdAmount) internal view returns (uint256) {
        uint256 ethPrice = getLatestEthPrice();
        // usdAmount (usd * 10^8) * 1e18 (wei/eth) / ethPrice (usd * 10^8) = wei
        return usdAmount.mul(1e18).div(ethPrice);
    }

    /**
     * @dev 获取ERC20代币到美元的价格
     * @param priceFeedAddress 代币/美元价格预言机地址
     * @return tokenPriceInUSD 代币的最新价格（单位：美元，乘以10^8）
     */
    function getLatestTokenPrice(address priceFeedAddress) internal view returns (uint256) {
        require(priceFeedAddress != address(0), 'Price feed address cannot be zero');
        AggregatorV3Interface priceFeed = AggregatorV3Interface(priceFeedAddress);
        (, int256 price, , , ) = priceFeed.latestRoundData();
        return uint256(price);
    }

    /**
     * @dev 将ERC20代币金额转换为美元
     * @param tokenAmount 代币金额
     * @param priceFeedAddress 代币/美元价格预言机地址
     * @return usdAmount 对应的美元金额（单位：美元，乘以10^8）
     */
    function convertTokenToUSD(uint256 tokenAmount, address priceFeedAddress) internal view returns (uint256) {
        uint256 tokenPrice = getLatestTokenPrice(priceFeedAddress);
        return tokenAmount.mul(tokenPrice);
    }
}