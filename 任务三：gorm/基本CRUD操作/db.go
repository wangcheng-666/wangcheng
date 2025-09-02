package config

import (
	"awesomeProject/models"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:woaiwd@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"

	// 配置 GORM 日志器
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到标准输出，带换行
		logger.Config{
			SlowThreshold:             time.Second, // 定义慢 SQL 阈值
			LogLevel:                  logger.Info, // 关键：设为 Info 级别，显示 SQL
			IgnoreRecordNotFoundError: true,        // 忽略记录未找到的错误（比如 Find 没查到）
			Colorful:                  true,        // 启用颜色（可选）
		},
	)

	// 打开数据库连接，使用自定义日志器
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, // 使用我们定义的日志器
	})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。

	DB = db
	DB.AutoMigrate(&models.Student{})
	student := &models.Student{}
	var students []models.Student
	student.Name = "张三"
	student.Age = 20
	student.Grade = "三年级"
	result := DB.Create(&student)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println("!111111")
	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	re := db.Where("age>?", 18).Find(&students)
	if re.Error != nil {
		fmt.Println(re.Error)
	} else {
		fmt.Printf("查到了：%d\n", re.RowsAffected)
		for _, val := range students {
			fmt.Printf("姓名:%s年龄:%d年级：%s\n", val.Name, val.Age, val.Grade)
		}
	}
	// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	db.Model(&models.Student{}).Where("name=?", "张三").Update("grade", "四年级")
	db.Where("grade=?", "四年级").Updates(&models.Student{
		Name: "小周222",
		Age:  11,
	})
	// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	db.Where("age<?", 15).Delete(&models.Student{})
}
