package models

// 账户信息表
type Account struct {
	Id      uint    `gorm:"primary_key;auto_increment"`
	Balance float64 `gorm:"type:decimal(10,2)"`
}

// 转账记录表
type Transactions struct {
	Id int `gorm:"primary_key"`
	//转出账户ID
	FromAccountId uint    `gorm:"index;not null;comment:转出账户ID"`
	toAccountId   uint    `gorm:"index;not null;comment:接受账户ID"`
	Amount        float64 `gorm:"type:decimal(10,2);comment:转账金额"`
}
