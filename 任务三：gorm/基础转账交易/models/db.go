package models

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// 检查账户余额
func (account *Account) Withdraw(amount float64, tx *gorm.DB) error {
	if account.Balance < amount {
		return errors.New("账户余额不足")
	}
	result := tx.Model(account).Update("balance", account.Balance-amount)
	return result.Error
}

// 转入操作
func (account *Account) Deposit(amount float64, tx *gorm.DB) error {
	result := tx.Model(account).Update("balance", account.Balance+amount)
	return result.Error
}

// 记录转账信息
func CreateTransation(amount float64, fromId uint, toId uint, tx *gorm.DB) error {
	tx.AutoMigrate(&Account{})
	tra := &Transactions{
		FromAccountId: fromId,
		Amount:        amount,
		toAccountId:   toId,
	}
	if err := tx.Create(tra).Error; err != nil {
		return err
	}
	return nil
}

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
	db.AutoMigrate(&Account{})
	account := &[]Account{
		{Balance: 100.85},
		{Balance: 0.01},
	}
	db.Create(account)
	Transfer(11, 12, 100)
}

// 执行事务
func Transfer(fromId uint, toId uint, amount float64) error {
	if amount <= 0 {
		return errors.New(fmt.Sprintf("余额%2f不够交易！！！", amount))
	}
	return DB.Transaction(func(tx *gorm.DB) error {
		//声明结构体把查到的数据放进去
		var fromAccount, toAccount *Account
		if err := tx.Where("id = ?", fromId).First(&fromAccount).Error; err != nil {
			return err
		}
		if fromAccount.Balance < amount {
			return errors.New("账户余额不足")
		}
		if err := tx.Where("id=?", toId).First(&toAccount).Error; err != nil {
			return errors.New("用户不存在")
		}
		if err := fromAccount.Withdraw(amount, tx); err != nil {
			return err
		}
		if err := toAccount.Deposit(amount, tx); err != nil {
			return err
		}
		err := CreateTransation(amount, fromAccount.Id, toAccount.Id, tx)
		if err != nil {
			return err
		}
		return nil
	})
}
