package main

import "fmt"

// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
// 组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("姓名: %s, 年龄: %d, 员工ID: %d\n", e.Person.name, e.Person.age, e.EmployeeID)
}

func main() {
	pers := &Person{name: "张三", age: 18}
	emp := &Employee{Person: *pers, EmployeeID: 1}
	emp.PrintInfo()
	fmt.Println(emp.name, emp.age, emp.EmployeeID)
}
