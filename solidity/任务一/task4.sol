// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract IntegerToRoman {
    function intToRoman(uint256 num) public pure returns (string memory) {
        if (num == 0) {
            return "";
        }

        uint256[13] memory values = [
            uint256(1000), uint256(900), uint256(500), uint256(400),
            uint256(100),  uint256(90),  uint256(50),  uint256(40),
            uint256(10),   uint256(9),   uint256(5),   uint256(4),
            uint256(1)
        ];

        string[13] memory numerals = [
            "M", "CM", "D", "CD",
            "C", "XC", "L", "XL",
            "X", "IX", "V", "IV", "I"
        ];

        bytes memory result = "";

        for (uint256 i = 0; i < 13; i++) {
            uint256 count = num / values[i];
            for (uint256 j = 0; j < count; j++) {
                result = bytes(abi.encodePacked(result, numerals[i]));
            }
            num -= values[i] * count;
        }

        return string(result);
    }
}