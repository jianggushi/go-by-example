package main

import "fmt"

func main() {
	messages := make(chan string) // channel 是 goroutine 可以进行通信的管道，可以通过 make() 方法定义

	go func() { messages <- "ping" }() // 通过 channel <- value 方法向通道中发送数据

	msg := <-messages // 通过 <- channel 从通道中接收数据
	fmt.Println(msg)  // 这里发送和接收都是阻塞的，所以先启动一个 goroutine 在其中发送数据，main routine 中阻塞等待接收数据
}
