package main

import (
	"fmt"
)

// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。
func GetWg(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}
func main() {
	ch := make(chan int, 100)
	go GetWg(ch)
	for value := range ch {
		fmt.Printf("数字%d被打印\n", value)
	}
}
