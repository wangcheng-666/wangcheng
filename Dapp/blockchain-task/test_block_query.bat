@echo off

REM 区块链读写任务 - 区块查询测试脚本
REM 此脚本不需要私钥配置，可以直接运行

echo 正在查询Sepolia测试网络的最新区块信息...
echo =======================================
go run cmd/block_query/main.go
echo =======================================
echo 区块查询完成。
echo 
echo 如果您想查询特定区块，请使用以下命令格式：
echo go run cmd/block_query/main.go 5000000
echo 
echo 例如：go run cmd/block_query/main.go 5000000
echo 
pause