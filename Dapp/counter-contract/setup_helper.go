package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	fmt.Println("===== 以太坊智能合约环境设置助手 =====")
	fmt.Println("这个工具将帮助您快速设置开发环境")
	fmt.Println("===================================")
	fmt.Println("提示：对于新手，推荐直接运行simple_start.bat一键脚本")
	fmt.Println("===================================")

	// 检查操作系统
	if runtime.GOOS != "windows" {
		fmt.Println("警告: 此工具为Windows系统优化。在其他系统上可能需要调整命令。")
	}

	// 创建必要的目录
	createDirIfNotExists("contracts")
	createDirIfNotExists("abi")
	createDirIfNotExists("bin")
	createDirIfNotExists("go-contract")

	// 显示菜单
	showMenu()
}

// 创建目录如果不存在
func createDirIfNotExists(dirName string) {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			fmt.Printf("创建目录 %s 失败: %v\n", dirName, err)
		} else {
			fmt.Printf("✓ 已创建目录: %s\n", dirName)
		}
	}
}

// 显示主菜单
func showMenu() {
	for {
		fmt.Println("\n请选择操作:")
		fmt.Println("1. 检查环境工具是否安装")
		fmt.Println("2. 编译智能合约")
		fmt.Println("3. 生成Go绑定代码")
		fmt.Println("4. 退出")

		var choice int
		fmt.Print("请输入选择 (1-4): ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			checkEnvironment()
		case 2:
			compileContract()
		case 3:
			generateBindings()
		case 4:
			fmt.Println("谢谢使用！再见。")
			return
		default:
			fmt.Println("无效的选择，请重新输入")
		}
	}
}

// 检查环境工具是否安装
func checkEnvironment() {
	fmt.Println("\n===== 检查环境工具 =====")

	// 检查solcjs
	checkTool("solcjs", "npm install -g solc", "Solidity编译器")

	// 检查abigen
	checkTool("abigen", "go install github.com/ethereum/go-ethereum/cmd/abigen@latest", "Go绑定生成工具")

	// 检查Go
	checkTool("go", "请从 https://go.dev/dl/ 安装Go语言", "Go语言")

	// 检查Node.js
	checkTool("node", "请从 https://nodejs.org/ 安装Node.js", "Node.js")

	fmt.Println("\n环境检查完成！")
}

// 检查特定工具是否安装
func checkTool(toolName, installCommand, toolDescription string) {
	cmd := exec.Command("where", toolName)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("✗ %s 未安装\n", toolDescription)
		fmt.Printf("  安装命令: %s\n", installCommand)
	} else {
		fmt.Printf("✓ %s 已安装\n", toolDescription)
	}
}

// 编译智能合约
func compileContract() {
	fmt.Println("\n===== 编译智能合约 =====")

	// 检查合约文件是否存在
	contractFile := "contracts\\Counter.sol"
	if _, err := os.Stat(contractFile); os.IsNotExist(err) {
		fmt.Printf("错误: 合约文件 %s 不存在\n", contractFile)
		fmt.Println("请先创建智能合约文件")
		return
	}

	// 编译合约
	fmt.Println("正在编译合约...")
	cmd := exec.Command("cmd", "/c", "solcjs", "--abi", "--bin", contractFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("编译失败: %v\n", err)
		fmt.Printf("错误输出: %s\n", string(output))
		return
	}

	// 移动生成的文件
	renameAndMoveFiles()

	fmt.Println("✓ 合约编译成功！")
}

// 重命名并移动编译生成的文件
func renameAndMoveFiles() {
	abiSource := "contracts_Counter_sol_Counter.abi"
	binSource := "contracts_Counter_sol_Counter.bin"
	abiDest := "abi\\Counter.abi"
	binDest := "bin\\Counter.bin"

	// 移动ABI文件
	if err := os.Rename(abiSource, abiDest); err != nil {
		fmt.Printf("移动ABI文件失败: %v\n", err)
	} else {
		fmt.Printf("✓ ABI文件: %s\n", abiDest)
	}

	// 移动字节码文件
	if err := os.Rename(binSource, binDest); err != nil {
		fmt.Printf("移动字节码文件失败: %v\n", err)
	} else {
		fmt.Printf("✓ 字节码文件: %s\n", binDest)
	}
}

// 生成Go绑定代码
func generateBindings() {
	fmt.Println("\n===== 生成Go绑定代码 =====")

	// 检查必要的文件是否存在
	abiFile := "abi\\Counter.abi"
	binFile := "bin\\Counter.bin"

	if _, err := os.Stat(abiFile); os.IsNotExist(err) {
		fmt.Printf("错误: ABI文件 %s 不存在\n", abiFile)
		fmt.Println("请先编译智能合约")
		return
	}

	if _, err := os.Stat(binFile); os.IsNotExist(err) {
		fmt.Printf("错误: 字节码文件 %s 不存在\n", binFile)
		fmt.Println("请先编译智能合约")
		return
	}

	// 生成Go绑定代码
	fmt.Println("正在生成Go绑定代码...")
	cmd := exec.Command("cmd", "/c", "abigen", "--abi", abiFile, "--bin", binFile, "--pkg", "counter", "--out", "go-contract\\counter.go")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("生成绑定代码失败: %v\n", err)
		fmt.Printf("错误输出: %s\n", string(output))
		return
	}

	fmt.Println("✓ Go绑定代码已生成到 go-contract\\counter.go")
}