// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract RomanToInteger {
    //     字符          数值
    // I             1
    // V             5
    // X             10
    // L             50
    // C             100
    // D             500
    // M             1000
    // 将罗马数字字符串转换为整数
    function romanToInt(string memory s) public pure returns (int256) {
        bytes memory strBytes = bytes(s);
        uint256 len = strBytes.length;
        int256 result = 0;

        for (uint256 i = 0; i < len; ) {
            bytes1 char = strBytes[i];

            // 检查双字符特殊情况
            if (i + 1 < len) {
                bytes1 nextChar = strBytes[i + 1];
                if (char == 'I' && nextChar == 'V') {
                    result += 4;
                    i += 2;
                    continue;
                } else if (char == 'I' && nextChar == 'X') {
                    result += 9;
                    i += 2;
                    continue;
                } else if (char == 'X' && nextChar == 'L') {
                    result += 40;
                    i += 2;
                    continue;
                } else if (char == 'X' && nextChar == 'C') {
                    result += 90;
                    i += 2;
                    continue;
                } else if (char == 'C' && nextChar == 'D') {
                    result += 400;
                    i += 2;
                    continue;
                } else if (char == 'C' && nextChar == 'M') {
                    result += 900;
                    i += 2;
                    continue;
                }
            }

            // 单字符情况
            if (char == 'I') {
                result += 1;
            } else if (char == 'V') {
                result += 5;
            } else if (char == 'X') {
                result += 10;
            } else if (char == 'L') {
                result += 50;
            } else if (char == 'C') {
                result += 100;
            } else if (char == 'D') {
                result += 500;
            } else if (char == 'M') {
                result += 1000;
            } else {
                revert("Invalid Roman Character");
            }
            i++;
        }
        return result;
    }
}