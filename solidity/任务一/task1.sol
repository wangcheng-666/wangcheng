// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract Voting {
    address public admin;

    constructor() {
        admin = msg.sender;
    }
    // 默认伦茨
    uint256 public currentRound = 1;

    // 地址、轮次、票数
    mapping(address => mapping(uint256 => uint256)) public votes;
    // 是否投票
    mapping(address => uint256) public lastVoteRound;

    // 定义自定义错误，好像是0.8才有的新特性更省gas
    error InvalidCandidate();
    error AlreadyVoted();
    error NotAdmin();

    // 投票给指定候选人
    function vote(address candidate) external {
        if (candidate == address(0)) {
            revert InvalidCandidate();
        }
        if (lastVoteRound[msg.sender] >= currentRound) {
            revert AlreadyVoted();
        }

        lastVoteRound[msg.sender] = currentRound;
        votes[candidate][currentRound] += 1;
    }

    // 获取候选人当前轮次票数
    function getVotes(address candidate) external view returns (uint256) {
        if (candidate == address(0)) {
            revert InvalidCandidate();
        }
        return votes[candidate][currentRound];
    }

    // 管理员获取指定轮次候选人的票数
    function AdminGetVote(address candidate, uint256 _currentRound) external view returns (uint256) {
        if (msg.sender != admin) {
            revert NotAdmin();
        }
        if (candidate == address(0)) {
            revert InvalidCandidate();
        }
        return votes[candidate][_currentRound];
    }

    // 重置投票：进入下一轮
    function resetVotes() external {
        currentRound++;
    }
}