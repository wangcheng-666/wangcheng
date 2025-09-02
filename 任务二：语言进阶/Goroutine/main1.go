package main

import (
	"fmt"
	"sync"
	"time"
)

func Singular(fh []int, ch chan int, group *sync.WaitGroup) {
	defer group.Done()
	for _, v := range fh {
		if v%2 != 0 {
			ch <- v
			fmt.Println("奇数", v)
			time.Sleep(1 * time.Second)
		}
	}
}
func EvenNumber(fh []int, ch chan int, group *sync.WaitGroup) {
	defer group.Done()
	for _, v := range fh {
		if v%2 != 1 {
			ch <- v
			fmt.Println("偶数", v)
			time.Sleep(1 * time.Second)
		}
	}
	close(ch)
}

func main() {
	fh := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go Singular(fh, ch, &wg)
	go EvenNumber(fh, ch, &wg)
	go func() {
		wg.Wait()
	}()
	for val := range ch {
		fmt.Println(val)
	}
}
