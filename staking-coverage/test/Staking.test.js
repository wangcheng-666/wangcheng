const { expect } = require("chai");
const { ethers } = require("hardhat");

// 正常质押流程
// 边界情况（零金额、超额质押）
// 多次质押逻辑
// 新用户初始化
// 错误处理（回滚）
// 多用户场景
// 事件触发
// 状态变化验证
describe("Staking合约完整测试", function () {
    let staking;
    let token;
    let owner, user1, user2;

    beforeEach(async function () {
        [owner, user1, user2] = await ethers.getSigners();
        
        const TokenFactory = await ethers.getContractFactory("ERC20Mock");
        token = await TokenFactory.deploy(ethers.parseEther("1000000"));
        await token.waitForDeployment();
        
        const Staking = await ethers.getContractFactory("Staking");
        staking = await Staking.deploy(await token.getAddress());
        await staking.waitForDeployment();
        
        await token.transfer(user1.address, ethers.parseEther("1000"));
        await token.transfer(user2.address, ethers.parseEther("1000"));
        
        await token.connect(user1).approve(await staking.getAddress(), ethers.parseEther("1000"));
        await token.connect(user2).approve(await staking.getAddress(), ethers.parseEther("1000"));
    });

    describe("stake 函数测试", function () {
        it("应该成功质押代币", async function () {
            const stakeAmount = ethers.parseEther("100");
            
            // 记录前置状态
            const userBalanceBefore = await token.balanceOf(user1.address);
            const contractBalanceBefore = await token.balanceOf(await staking.getAddress());
            const stakedBalanceBefore = await staking.stakedBalance(user1.address);
            const totalStakedBefore = await staking.totalStaked();
            
            // 执行质押
            await expect(staking.connect(user1).stake(stakeAmount))
                .to.emit(staking, "Staked")
                .withArgs(user1.address, stakeAmount);
            
            // 验证后置状态
            const userBalanceAfter = await token.balanceOf(user1.address);
            const contractBalanceAfter = await token.balanceOf(await staking.getAddress());
            const stakedBalanceAfter = await staking.stakedBalance(user1.address);
            const totalStakedAfter = await staking.totalStaked();
            const stakingTime = await staking.stakingStarTtime(user1.address);
            
            // 验证代币转移
            expect(userBalanceAfter).to.equal(userBalanceBefore - stakeAmount);
            expect(contractBalanceAfter).to.equal(contractBalanceBefore + stakeAmount);
            
            // 验证质押记录
            expect(stakedBalanceAfter).to.equal(stakeAmount);
            expect(stakedBalanceAfter).to.equal(stakedBalanceBefore + stakeAmount);
            
            // 验证总质押量
            expect(totalStakedAfter).to.equal(totalStakedBefore + stakeAmount);
            
            // 验证质押时间设置
            expect(stakingTime).to.be.gt(0);
            
            console.log("质押测试通过");
        });

        it("应该拒绝零金额质押", async function () {
            await expect(staking.connect(user1).stake(0))
                .to.be.revertedWith("Amount must be greater than 0");
        });

        it("应该拒绝超过余额的质押", async function () {
            const excessiveAmount = ethers.parseEther("2000");
            await expect(staking.connect(user1).stake(excessiveAmount))
                .to.be.revertedWith("Insufficient balance");
        });

        it("应该处理多次质押", async function () {
            const firstStake = ethers.parseEther("50");
            const secondStake = ethers.parseEther("30");
            
            // 第一次质押
            await staking.connect(user1).stake(firstStake);
            const firstStakingTime = await staking.stakingStarTtime(user1.address);
            
            // 等待一段时间
            await ethers.provider.send("evm_increaseTime", [300]); // 5分钟
            await ethers.provider.send("evm_mine");
            
            // 第二次质押
            await staking.connect(user1).stake(secondStake);
            
            // 验证总质押金额
            const totalStaked = await staking.stakedBalance(user1.address);
            expect(totalStaked).to.equal(firstStake + secondStake);
            
            // 验证质押时间更新
            const secondStakingTime = await staking.stakingStarTtime(user1.address);
            expect(secondStakingTime).to.be.gt(firstStakingTime);
            
            // 验证累积了奖励
            const rewards = await staking.rewards(user1.address);
            expect(rewards).to.be.gt(0);
        });

        it("新用户质押应该设置初始时间", async function () {
            const stakeAmount = ethers.parseEther("100");
            
            // 验证初始时间为0
            const initialTime = await staking.stakingStarTtime(user1.address);
            expect(initialTime).to.equal(0);
            
            // 执行质押
            await staking.connect(user1).stake(stakeAmount);
            
            // 验证时间已设置
            const stakingTime = await staking.stakingStarTtime(user1.address);
            expect(stakingTime).to.be.gt(0);
        });

        it("质押失败时应该回滚所有状态变化", async function () {
            const stakeAmount = ethers.parseEther("100");
            const initialBalance = await token.balanceOf(user1.address);
            
            // 尝试质押但授权不足（设置低授权额度）
            await token.connect(user1).approve(await staking.getAddress(), ethers.parseEther("50"));
            
            await expect(staking.connect(user1).stake(stakeAmount))
                .to.be.reverted; // transferFrom 会失败
            
            // 验证状态没有改变
            const finalBalance = await token.balanceOf(user1.address);
            const stakedBalance = await staking.stakedBalance(user1.address);
            const totalStaked = await staking.totalStaked();
            
            expect(finalBalance).to.equal(initialBalance);
            expect(stakedBalance).to.equal(0);
            expect(totalStaked).to.equal(0);
        });

        it("多个用户可以同时质押", async function () {
            const user1Stake = ethers.parseEther("100");
            const user2Stake = ethers.parseEther("200");
            
            // 用户1质押
            await staking.connect(user1).stake(user1Stake);
            
            // 用户2质押
            await staking.connect(user2).stake(user2Stake);
            
            // 验证各自的质押余额
            const user1Balance = await staking.stakedBalance(user1.address);
            const user2Balance = await staking.stakedBalance(user2.address);
            const totalStaked = await staking.totalStaked();
            
            expect(user1Balance).to.equal(user1Stake);
            expect(user2Balance).to.equal(user2Stake);
            expect(totalStaked).to.equal(user1Stake + user2Stake);
            
            // 验证各自的质押时间
            const user1Time = await staking.stakingStarTtime(user1.address);
            const user2Time = await staking.stakingStarTtime(user2.address);
            
            expect(user1Time).to.be.gt(0);
            expect(user2Time).to.be.gt(0);
            expect(user1Time).to.not.equal(user2Time);
        });
    });

    describe("取消质押测试", function () {
        beforeEach(async function () {
            // 用户先质押100个代币
            await staking.connect(user1).stake(ethers.parseEther("100"));
        });

        it("用户应该能够取消质押", async function () {
            const unstakeAmount = ethers.parseEther("50");
            
            // 记录取消质押前的余额
            const userBalanceBefore = await token.balanceOf(user1.address);
            const contractBalanceBefore = await token.balanceOf(await staking.getAddress());
            
            // 用户取消质押
            await staking.connect(user1).unstake(unstakeAmount);
            
            // 检查质押余额是否正确减少
            const stakedBalance = await staking.stakedBalance(user1.address);
            expect(stakedBalance).to.equal(ethers.parseEther("50"));
            
            // 检查总质押量是否正确减少
            const totalStaked = await staking.totalStaked();
            expect(totalStaked).to.equal(ethers.parseEther("50"));
            
            // 检查代币是否实际返还给用户
            const userBalanceAfter = await token.balanceOf(user1.address);
            const contractBalanceAfter = await token.balanceOf(await staking.getAddress());
            
            expect(userBalanceAfter).to.equal(userBalanceBefore + unstakeAmount);
            expect(contractBalanceAfter).to.equal(contractBalanceBefore - unstakeAmount);
            
            console.log("取消质押测试通过");
        });

        it("应该拒绝超过质押余额的取消质押", async function () {
            const excessiveAmount = ethers.parseEther("200");
            await expect(
                staking.connect(user1).unstake(excessiveAmount)
            ).to.be.revertedWith("Insufficient staked balance");
            console.log("超额取消质押测试通过");
        });
    });

    describe("奖励计算测试", function () {
        beforeEach(async function () {
            await staking.connect(user1).stake(ethers.parseEther("100"));
        });

        it("应该正确计算奖励", async function () {
            //  以分钟为单位，增加61秒确保至少有1分钟
            await ethers.provider.send("evm_increaseTime", [61]);
            await ethers.provider.send("evm_mine");
            
            const reward = await staking.calculateReward(user1.address);
            expect(reward).to.be.gt(0);
            console.log("1分钟奖励:", reward.toString());
        });

        it("不足1分钟应该没有奖励", async function () {
            //  只增加59秒，不足1分钟
            await ethers.provider.send("evm_increaseTime", [59]);
            await ethers.provider.send("evm_mine");
            
            const reward = await staking.calculateReward(user1.address);
            expect(reward).to.equal(0);
            console.log("59秒奖励:", reward.toString());
        });
    });

    describe("领取奖励测试", function () {
        beforeEach(async function () {
            // 用户质押
            await staking.connect(user1).stake(ethers.parseEther("100"));
            
            //  等待足够的时间（比如5分钟）来获得明显的奖励
            await ethers.provider.send("evm_increaseTime", [300]); // 5分钟
            await ethers.provider.send("evm_mine");
            
            // 给质押合约充值奖励代币
            await token.transfer(await staking.getAddress(), ethers.parseEther("1000"));
        });

        it("用户应该能够领取奖励", async function () {
            const userBalanceBefore = await token.balanceOf(user1.address);
            
            // 计算预期奖励
            const expectedReward = await staking.getTotalRewards(user1.address);
            expect(expectedReward).to.be.gt(0);
            
            // 用户领取奖励
            await staking.connect(user1).claimReward();
            
            const userBalanceAfter = await token.balanceOf(user1.address);
            const actualReward = userBalanceAfter - userBalanceBefore;
            
            //  现在可以精确比较，因为以分钟为单位消除了秒级差异
            expect(actualReward).to.equal(expectedReward);
            
            console.log("领取奖励测试通过，奖励:", actualReward.toString());
        });
    });

    describe("事件触发测试", function () {
        it("领取奖励应该触发RewardClaimed事件", async function () {
            await staking.connect(user1).stake(ethers.parseEther("100"));
            //  等待足够时间
            await ethers.provider.send("evm_increaseTime", [300]); // 5分钟
            await ethers.provider.send("evm_mine");
            
            await token.transfer(await staking.getAddress(), ethers.parseEther("1000"));
            
            const expectedReward = await staking.getTotalRewards(user1.address);
            
            //  现在可以精确匹配事件参数
            await expect(staking.connect(user1).claimReward())
                .to.emit(staking, "RewardClaimed")
                .withArgs(user1.address, expectedReward);
            
            console.log("RewardClaimed事件触发测试通过");
        });
    });
});