package main

import (
	"fmt"
	"sync"
)

// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
func GetWg(wg *sync.WaitGroup, mu *sync.Mutex, count *int) {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		mu.Lock()
		fmt.Printf("1111%d,总值为%d\n", i, *count)
		*count++
		mu.Unlock()
	}
	fmt.Printf("2222%d\n", *count)
}
func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	var count = 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go GetWg(&wg, &mu, &count)
	}
	wg.Wait()
	fmt.Printf("3333%d\n", count)
	//counts := 0
	//for v := range ch {
	//	counts += v

	//}
}
