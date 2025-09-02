package main

import (
	"fmt"
	"sync"
	"time"
)

type task func()

// 任务调度器
func TaskScheduler(tasks []task) {
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for i, Task := range tasks {
		go func(tkNumber int, tk task) {
			tim := time.Now()
			defer wg.Done()
			tk()
			timeOut := time.Since(tim)
			fmt.Printf("第%d个方法，执行时间为：%v\n", i+1, timeOut)
		}(i, Task)
	}
	wg.Wait()
}

func main() {
	tm := time.Now()
	tasks := []task{
		func() {
			time.Sleep(12 * time.Millisecond)
			fmt.Println("任务A执行时间为：", time.Since(tm))
		},
		func() {
			time.Sleep(123 * time.Millisecond)
			fmt.Println("任务B执行时间为：", time.Since(tm))
		},
		func() {
			time.Sleep(53 * time.Millisecond)
			fmt.Println("任务C执行时间为：", time.Since(tm))
		},
	}
	TaskScheduler(tasks)
	fmt.Printf("总共执行时间%v\n", time.Since(tm))
}
