package main

import "fmt"

func ping(pings chan<- string, msg string) { // 函数参数中的 channel 可以限定方向，增加类型安全，chan<- 只能向 channel 发送消息
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) { // <-chan 只能从 channel 中接收消息，chan<- 只能向 channel 中发送消息
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
