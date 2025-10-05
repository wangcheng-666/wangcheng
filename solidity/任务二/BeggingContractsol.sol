// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";

contract BeggingContractsol is Ownable {
    mapping (address=>uint256) public donations;

    address[] public donors;
    uint256 public startTime;
    uint256 public endTime;

    event Downation(address indexed donor, uint256 amount);

    constructor(uint256 _durationInDays) Ownable(msg.sender){
        startTime = block.timestamp;
        endTime = block.timestamp+(_durationInDays * 1 days);
    }

    function donate()external payable {
        require(block.timestamp >= startTime,"Donation period has not started");
        require(block.timestamp < endTime,"Donation period has ended");
        require(msg.value > 0,"You must send some ETH");
        if (donations[msg.sender] == 0){
            donors.push(msg.sender); //新用户添加
        }
        donations[msg.sender] += msg.value;
        emit Downation(msg.sender,msg.value);
    }

    function withdraw()external onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0,"No funds to withdraw");
        (bool sent, )= payable(owner()).call{value:balance}("");
        require(sent,"Withdrawal failed");
    }

    function getDonation(address _donor)external view returns (uint256) {
        return donations[_donor];
    }

    function getDonorCount()external view returns (uint256) {
        return donors.length;
    }

    function getTop3Donors() external view returns (address[3] memory, uint256[3] memory) {
        uint256[3] memory amounts;
        address[3] memory top3Addresses;

        for (uint256 i = 0; i < donors.length; ++i) {
            address donor = donors[i];
            uint256 amount = donations[donor];

            if (amount > amounts[0]) {
                amounts[2] = amounts[1];
                amounts[1] = amounts[0];
                amounts[0] = amount;
                top3Addresses[2] = top3Addresses[1];
                top3Addresses[1] = top3Addresses[0];
                top3Addresses[0] = donor;
            } else if (amount > amounts[1]) {
                amounts[2] = amounts[1];
                amounts[1] = amount;

                top3Addresses[2] = top3Addresses[1];
                top3Addresses[1] = donor;
            } else if (amount > amounts[2]) {
                amounts[2] = amount;
                top3Addresses[2] = donor;
            }
        }

        return (top3Addresses, amounts);
    }

    function isDonationOpen() external view returns (bool) {
        return (block.timestamp >= startTime && block.timestamp < endTime);
    }
}