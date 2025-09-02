// main.go
package main

import (
	"awesomeProject/models"
	"log"
)

func main() {
	// 初始化数据库连接
	models.ConnectDatabase()

	// 检查是否连接成功
	sqlDB, err := models.DB.DB()
	if err != nil {
		log.Fatal("获取 DB 实例失败:", err)
	}
	err = sqlDB.Ping()

	if err != nil {
		log.Fatal("数据库 Ping 失败:", err)
	}

	log.Println("✅ 数据库连接成功！")

	// 启动服务、路由、业务逻辑...
}
