package models

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDatabase() {
	dsn := "root:woaiwd@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	// 使用 sqlx 连接数据库
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	// 可选：设置连接池
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	DB = db
	result, _ := DB.Exec("insert into employees(Name,Department,Salary) values (?,?,?)", "张三", "技术部", 293.85)
	fmt.Println(result.LastInsertId())
	//Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	var empl []employees
	res := DB.Select(&empl, "select * from employees where Department like ?", "%技术部%")
	if res != nil {
		log.Fatal("查询失败:", res)
	}
	for _, re := range empl {
		fmt.Println(re)
	}
	//使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	var sMax employees
	maxErr := DB.Get(&sMax, "SELECT *  from employees ORDER BY salary  desc LIMIT  1")
	if maxErr != nil {
		log.Fatal(maxErr)
	}
	fmt.Printf("姓名%s,最高工资：%2f", sMax.Name, sMax.Salary)
}
