package main

import (
	"fmt"
)

func Getstr(arrInt []int, number int) []int {
	m := make(map[int]int)
	for i := 0; i < len(arrInt); i++ {
		next := number - arrInt[i]

		if j, ok := m[next]; ok {
			return []int{j, i}
		}
		m[arrInt[i]] = i
	}
	return []int(nil)
}

func main() {
	arrInt := []int{5, 7, 3, 1, 2, 3, 5, 10}
	number := 15
	k := Getstr(arrInt, number)
	sprintf := fmt.Sprintf("总数为：%d\n下标结果：%d", number, k)
	fmt.Println(sprintf)
}
