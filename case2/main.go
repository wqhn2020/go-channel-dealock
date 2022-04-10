package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义一个WainGroup，用于实现优雅退出
var wg sync.WaitGroup

// produce 每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
func produce(queue chan<- int) {
	go func() {
		for {
			queue <- 0 // 把0放入通道。如果通道满了，则这里会阻塞。
			fmt.Printf("放入一个元素\n")
			time.Sleep(time.Second)
		}
	}()
}

func consume(queue <-chan int) {
	go func() {
		for {
			<-queue
			fmt.Printf("--> 取出一个元素\n")
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	wg.Add(2) // 设置等待子协程退出

	queue := make(chan int, 3)
	go produce(queue)

	go consume(queue) // 启动消费者协程

	time.Sleep(time.Second * 5)
	fmt.Println("case2 waiting")
	wg.Wait() // 等待子协程退出，这种方式比较优雅。
	fmt.Println("-- done --")
}

