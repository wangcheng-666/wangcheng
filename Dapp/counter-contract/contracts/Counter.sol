// SPDX-License-Identifier: MIT
// Counter.sol - 以太坊计数器合约
pragma solidity ^0.8.0;

// 定义事件，当计数增加时触发
event CountIncreased(uint256 indexed newValue);

contract Counter {
    // 声明一个公共状态变量用于存储计数
    uint256 public count = 0;
    
    // increase函数 - 用于增加计数
    function increase() public {
        count += 1;
        emit CountIncreased(count); // 触发事件
    }
    
    // getCount函数 - 用于获取当前计数
    function getCount() public view returns (uint256) {
        return count;
    }
}