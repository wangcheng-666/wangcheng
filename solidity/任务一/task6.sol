// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract SearchInSortedArray {
    uint256[5] public instart = [1, 2, 3, 6, 9];

    // 二分查找：返回是否找到目标值
    function binarySearch(uint256 target) external view returns (bool) {
        uint256 left = 0;
        uint256 rigin = instart.length;
        while (left< rigin) 
        {
            uint256 min = left +(rigin-left )/2;
            if(instart[min] == target){
                return true;
            }else if(instart[min]>target){
                rigin = min--;
            }else if(instart[min]<target) {
                left = min++;
            }
        }
        return false;
    }
}