package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string) // 启动两个 goroutine 向两个 channel 中发送消息
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ { // 通过 select 等待多个 channel 中是否有消息到达
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
