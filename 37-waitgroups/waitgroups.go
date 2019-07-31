package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
	wg.Done() // worker 完成，调用 wg.Done() 通知已完成
}

func main() {
	var wg sync.WaitGroup // 定义一个 WaitGroup

	for i := 1; i <= 5; i++ { // 启动 5 个 worker，每启动一个 worker，wg 计数增加 1
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait() // main routine 调用 wg.Wait() 等待所有 worker 完成
}
