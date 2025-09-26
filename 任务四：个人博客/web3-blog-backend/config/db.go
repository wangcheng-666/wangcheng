package config

import (
	"web3-blog-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	var err error
	sdn := "root:woaiwd@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	if err != nil {
		panic("failed to connect to MySQL database: " + err.Error())
	}
	DB, err = gorm.Open(mysql.Open(sdn), &gorm.Config{
		// ✅ 开启日志：打印 SQL
		Logger: logger.Default.LogMode(logger.Info), // 会打印所有 SQL
	})
	// 自动迁移表结构（创建表）
	DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}
