// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract Voting {

    string public resultStr;

    function rangeStr(string memory str) external {
        bytes memory strBytes=  bytes(str);
        bytes memory result = new bytes(strBytes.length);
        for (uint i = 0; i < strBytes.length; i++) 
        {
            result[i] = strBytes[strBytes.length -1 -i];
        }
        resultStr = string(result);
    }
}