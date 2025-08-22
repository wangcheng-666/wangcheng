package main

import "fmt"

func Getstr(arrInt []int) []int {
	mas := make(map[int]int)
	var s []int
	for _, i := range arrInt {
		mas[i]++
	}
	fmt.Println(mas)
	for key, d := range mas {
		if d == 1 {
			s = append(s, key)
		}
	}
	fmt.Println(s)
	return s
}

func main() {
	arrInt := []int{1, 1, 2, 3, 4, 5, 6, 7, 7}
	result := Getstr(arrInt)
	fmt.Println("数组长度为：", len(result), result)
}
