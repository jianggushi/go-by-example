package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) { // 使用 channel 来使得两个 goroutine 进行同步
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true // worker 执行完成后，向 channel 中发送消息，通知 main routine 已完成
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done // main routine 启动 worker 后，等待接收 channel 中的消息
}
