// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract MergeSortedArrays {
    function merge(int256[] memory arr1, int256[] memory arr2) 
        public 
        pure 
        returns (int256[] memory) 
    {
        uint256 len1 = arr1.length;
        uint256 len2 = arr2.length;
        uint256 totalLen = len1 + len2;

        // 创建结果数组（在 memory 中）
        int256[] memory result = new int256[](totalLen);

        uint256 i = 0; // arr1 的指针
        uint256 j = 0; // arr2 的指针
        uint256 k = 0; // result 的指针

        // 双指针合并：从小到大比较
        while (i < len1 && j < len2) {
            if (arr1[i] <= arr2[j]) {
                result[k] = arr1[i];
                i++;
            } else {
                result[k] = arr2[j];
                j++;
            }
            k++;
        }

        //长度不一样剩余得给补全不需要再对比了
        while (i < len1) {
            result[k] = arr1[i];
            i++;
            k++;
        }

        while (j < len2) {
            result[k] = arr2[j];
            j++;
            k++;
        }

        return result;
    }
}