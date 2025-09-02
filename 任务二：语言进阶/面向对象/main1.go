package main

import (
	"fmt"
	"math"
)

// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	Width, Height float64
}

func (pr *Rectangle) Area() float64 {
	return pr.Width * pr.Height
}
func (pr *Rectangle) Perimeter() float64 {
	return 2*pr.Width + 2*pr.Height
}

type Circle struct {
	Radius float64
}

func (ci *Circle) Area() float64 {
	return math.Pi * ci.Radius * ci.Radius
}

func (ci *Circle) Perimeter() float64 {
	return 2 * math.Pi * ci.Radius
}

func main() {
	sd := &Rectangle{Width: 15 - 8, Height: 25.6}
	fmt.Println(sd.Area())
	fmt.Println(sd.Perimeter())
	dd := &Circle{Radius: 50.9}
	fmt.Println(dd.Area())
	fmt.Println(dd.Perimeter())
}
