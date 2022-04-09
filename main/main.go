package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义一个WainGroup，用于实现优雅退出
var wg sync.WaitGroup

// worker 每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
func worker(queue chan<- int) {
	go func() {
		for {
			queue <- 0 // 把0放入通道。如果通道满了，则这里会阻塞。
			fmt.Printf("放入一个元素\n")
			time.Sleep(time.Second)
		}
		wg.Done() // 通知main协程，子协程已退出
	}()
}

func main() {
	wg.Add(1) // 设置等待一个协程退出

	queue := make(chan int, 3)
	go worker(queue)

	time.Sleep(time.Second * 5)
	wg.Wait() // 等待子协程退出，这种方式比较优雅。
	fmt.Println("-- done --")
}
