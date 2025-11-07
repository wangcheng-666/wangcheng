// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

// 引入ERC20合约
interface IERC20 {
    function transfer(address to, uint256 amount) external returns (bool);
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) external returns (bool);
    function balanceOf(address account) external view returns (uint256);
}

//质押，用户存入代币获取奖励
contract Staking {
    IERC20 public stakingToken;

    constructor(address _stakingToken) {
        stakingToken = IERC20(_stakingToken);
    }
    // 记录用户存入的代币数量
    mapping(address => uint256) public stakedBalance;

    // 记录每个地址开始质押的时间
    mapping(address => uint256) public stakingStarTtime;

    // 记录用户的累计的奖励
    mapping(address => uint256) public rewards;

    //总质押数量
    uint256 public totalStaked;

    //设置每秒代币获得多少数量
    uint256 public constant REWARD_RATE = 100;

    // 用户质押时触发
    event Staked(address indexed user, uint256 amount);

    //用户取消时触发
    event Unstaked(address indexed user, uint256 amount);

    //用户领取奖励的时候触发
    event RewardClaimed(address indexed user, uint256 reward);

    function stake(uint256 amount) external {
        require(amount > 0, "Amount must be greater than 0");
        require(
            stakingToken.balanceOf(msg.sender) >= amount,
            "Insufficient balance"
        );
        //如果有旧的质押就累加
        if (stakedBalance[msg.sender] > 0) {
            rewards[msg.sender] += calculateReward(msg.sender);
        }
        require(
            stakingToken.transferFrom(msg.sender, address(this), amount),
            "stake transfer failed"
        );
        // 更新用户质押余额
        stakedBalance[msg.sender] += amount;
        // 记录质押时间，如果是第一次质押
        stakingStarTtime[msg.sender] = block.timestamp;
        // 更新总质押量
        totalStaked += amount;
        // 触发质押事件
        emit Staked(msg.sender, amount);
    }

    // 取消质押金额
    function unstake(uint256 amount) external {
        require(amount > 0, "Amount must be greater than 0");
        require(
            stakedBalance[msg.sender] >= amount,
            "Insufficient staked balance"
        );
        // 先看累计了多少奖励
        rewards[msg.sender] += calculateReward(msg.sender);
        stakedBalance[msg.sender] -= amount;
        stakingStarTtime[msg.sender] = block.timestamp;
        totalStaked -= amount;
        require(
            stakingToken.transfer(msg.sender, amount),
            "Unstake transfer failed"
        );
        emit Unstaked(msg.sender, amount);
    }

    function claimReward() external {
        // 计算总奖励=待计算奖励+已累计的奖励
        uint256 reward = calculateReward(msg.sender) + rewards[msg.sender];
        // 检查
        require(reward > 0, "No rewards to claim");
        // 已领取后重置时间和奖励金额
        rewards[msg.sender] = 0;
        stakingStarTtime[msg.sender] = block.timestamp;
        // 奖励
        require(
            stakingToken.transfer(msg.sender, reward),
            "Reward transfer failed"
        );
        emit RewardClaimed(msg.sender, reward);
    }

    function calculateReward(address user) public view returns (uint256) {
        if (stakedBalance[user] == 0) {
            return 0;
        }
        uint256 stakingDuration = block.timestamp - stakingStarTtime[user];

        // 使用分钟或小时为单位，减少精度敏感性(不这样测试会先调用)
        uint256 durationInMinutes = stakingDuration / 60;
        if (durationInMinutes == 0) {
            return 0;
        }
        // 计算奖励：质押金额 * 持续时间(分钟) * 奖励率
        return (stakedBalance[user] * durationInMinutes * REWARD_RATE) / 1e18;
    }

    // 获取总奖励函数：返回用户的总奖励（待计算 + 已累积）
    function getTotalRewards(address user) public view returns (uint256) {
        return calculateReward(user) + rewards[user];
    }
}
