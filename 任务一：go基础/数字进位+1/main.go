package main

import "fmt"

func Getstr(arrInt []int, number int) []int {
	for i := len(arrInt) - 1; i >= 0; i-- {
		fmt.Println(i)
		arrInt[i] += number
		if arrInt[i] < 10 {
			return arrInt
		}
		arrInt[i] = 0
	}
	result := make([]int, len(arrInt)+1)
	result[0] = 1
	return result
}

func main() {
	arrInt := []int{9, 9, 9}
	number := 1
	result := Getstr(arrInt, number)
	fmt.Println(result)
}
