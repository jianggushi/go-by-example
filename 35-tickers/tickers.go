package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond) // time.NewTicker() 会创建一个 Ticker

	go func() {
		for t := range ticker.C { // 每隔一段指定的时间，会发送一次消息，可以用来进行间隔执行
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
