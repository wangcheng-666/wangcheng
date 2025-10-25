@echo off

REM 区块链读写任务 - 交易发送测试脚本

REM 检查私钥是否已配置
for /f "tokens=*" %%i in ('type .env ^| findstr /i "PRIVATE_KEY="') do set PRIVATE_KEY_LINE=%%i
echo %PRIVATE_KEY_LINE% | findstr /i "PRIVATE_KEY=" >nul
if errorlevel 1 (
echo 错误: 未找到PRIVATE_KEY配置！
echo 请先在.env文件中设置您的以太坊私钥。
echo 详细说明请查看GETTING_STARTED.md文件。
pause
exit /b 1
)

echo %PRIVATE_KEY_LINE% | findstr /i "your_actual_private_key_here" >nul
if not errorlevel 1 (
echo 错误: 私钥未设置为实际值！
echo 请在.env文件中替换PRIVATE_KEY的值为您的真实私钥。
echo 详细说明请查看GETTING_STARTED.md文件。
pause
exit /b 1
)

REM 检查接收地址是否已配置
for /f "tokens=*" %%i in ('type .env ^| findstr /i "RECIPIENT_ADDRESS="') do set RECIPIENT_LINE=%%i
echo %RECIPIENT_LINE% | findstr /i "RECIPIENT_ADDRESS=" >nul
if errorlevel 1 (
echo 错误: 未找到RECIPIENT_ADDRESS配置！
echo 请先在.env文件中设置接收地址。
echo 详细说明请查看GETTING_STARTED.md文件。
pause
exit /b 1
)

echo %RECIPIENT_LINE% | findstr /i "your_recipient_address_here" >nul
if not errorlevel 1 (
echo 错误: 接收地址未设置为实际值！
echo 请在.env文件中替换RECIPIENT_ADDRESS的值为您要发送到的地址。
echo 详细说明请查看GETTING_STARTED.md文件。
pause
exit /b 1
)

REM 提示用户
cls
echo ==================================================
echo 区块链读写任务 - 交易发送功能
 echo ==================================================
echo 警告：运行此脚本将使用您的私钥发送真实交易到Sepolia测试网络。
echo 请确保：
echo 1. 您的账户中有足够的Sepolia测试以太币
 echo 2. 您已在.env文件中正确配置了私钥和接收地址
 echo 3. 您了解私钥的重要性并已安全保管
 echo ==================================================

echo. 
echo 默认转账金额为1个测试以太币(10^18 Wei)。
echo 如果要指定其他金额，请使用以下命令格式：
echo test_send_transaction.bat 100000000000000000
echo 例如：test_send_transaction.bat 100000000000000000 (转账0.5 ETH)

echo. 
echo 按任意键继续，按Ctrl+C取消...
pause >nul

REM 执行交易发送命令
if "%1"=="" (
    go run cmd/send_transaction/main.go
) else (
    go run cmd/send_transaction/main.go %1
)

pause