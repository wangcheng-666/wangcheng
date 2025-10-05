// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract Voting {
    address public admin;

    constructor() {
        admin = msg.sender;
    }
    // Ĭ���״�
    uint256 public currentRound = 1;

    // ��ַ���ִΡ�Ʊ��
    mapping(address => mapping(uint256 => uint256)) public votes;
    // �Ƿ�ͶƱ
    mapping(address => uint256) public lastVoteRound;

    // �����Զ�����󣬺�����0.8���е������Ը�ʡgas
    error InvalidCandidate();
    error AlreadyVoted();
    error NotAdmin();

    // ͶƱ��ָ����ѡ��
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

    // ��ȡ��ѡ�˵�ǰ�ִ�Ʊ��
    function getVotes(address candidate) external view returns (uint256) {
        if (candidate == address(0)) {
            revert InvalidCandidate();
        }
        return votes[candidate][currentRound];
    }

    // ����Ա��ȡָ���ִκ�ѡ�˵�Ʊ��
    function AdminGetVote(address candidate, uint256 _currentRound) external view returns (uint256) {
        if (msg.sender != admin) {
            revert NotAdmin();
        }
        if (candidate == address(0)) {
            revert InvalidCandidate();
        }
        return votes[candidate][_currentRound];
    }

    // ����ͶƱ��������һ��
    function resetVotes() external {
        currentRound++;
    }
}