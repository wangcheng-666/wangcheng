const { ethers } = require("hardhat");

async function main() {
  console.log("开始部署质押合约系统...");

  // 获取部署者账户
  const [deployer] = await ethers.getSigners();
  console.log("部署者地址:", deployer.address);
  console.log("部署者余额:", ethers.utils.formatEther(await deployer.getBalance()), "ETH");

  // 1. 部署模拟DAI代币 (ERC20Mock)
  console.log("\n正在部署模拟DAI代币...");
  const TokenFactory = await ethers.getContractFactory("ERC20Mock");
  const token = await TokenFactory.deploy(ethers.utils.parseEther("1000000")); // 100万代币
  await token.deployed();
  console.log("模拟DAI部署成功!");
  console.log("  合约地址:", token.address);
  console.log("  总供应量: 1,000,000 代币");

  // 2. 部署Staking质押合约
  console.log("\n正在部署Staking质押合约...");
  const Staking = await ethers.getContractFactory("Staking");
  const staking = await Staking.deploy(token.address);
  await staking.deployed();
  console.log("Staking合约部署成功!");
  console.log("  合约地址:", staking.address);
  console.log("  质押代币:", token.address);

  // 3. 给测试账户分配代币（可选）
  console.log("\n 给测试账户分配代币...");
  const [_, user1, user2] = await ethers.getSigners();
  
  await token.transfer(user1.address, ethers.utils.parseEther("1000"));
  await token.transfer(user2.address, ethers.utils.parseEther("1000"));
  console.log("代币分配完成!");
  console.log("  User1 余额:", ethers.utils.formatEther(await token.balanceOf(user1.address)), "代币");
  console.log("  User2 余额:", ethers.utils.formatEther(await token.balanceOf(user2.address)), "代币");

  // 4. 授权Staking合约使用代币（可选）
  console.log("\n 设置授权...");
  const tokenWithUser1 = token.connect(user1);
  const tokenWithUser2 = token.connect(user2);
  
  await tokenWithUser1.approve(staking.address, ethers.utils.parseEther("1000"));
  await tokenWithUser2.approve(staking.address, ethers.utils.parseEther("1000"));
  console.log("授权设置完成!");

  console.log("\n 所有合约部署完成!");
  console.log("==========================================");
  console.log("部署摘要:");
  console.log("  模拟DAI地址:", token.address);
  console.log("  Staking合约地址:", staking.address);
  console.log("  部署者地址:", deployer.address);
  console.log("==========================================");
  
  // 返回合约实例，方便其他脚本使用
  return { token, staking, deployer };
}

// 执行部署
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error("部署失败:", error);
    process.exit(1);
  });