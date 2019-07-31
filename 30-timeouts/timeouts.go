package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select { // 可以通过 select 实现 timeout 的功能
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second): // time.After() 等待指定的时间后通过 channel 返回一个超时消息
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 3")
	}
}
