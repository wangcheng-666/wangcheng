package main

import (
	"fmt"
)

func CountMinMap(minMap [9]int) []int {
	mapCount := make(map[string]int)
	count := make(map[int]int)
	var slice []int
	for _, d := range minMap {
		count[d]++
	}
	fmt.Println(count)
	for s, e := range count {
		if e == 1 {
			slice = append(slice, s)
		}
		p := fmt.Sprintf("数字%d出现的次数为:%d", s, e)
		fmt.Println(p)
		mapCount[p] = e
	}
	return slice
}

func main() {
	minMap := [9]int{1, 4, 9, 2, 3, 3, 4, 1, 1}
	mp := CountMinMap(minMap)
	fmt.Println("返回：", mp)
}
