package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second) // timer.NewTimer() 会创建一个定时器

	<-timer1.C // 定时器超时后，会通过定时器的 channel 发送消息，经常用来指定在多长时间后定时执行
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop() // 定时器在超时之前也可以被停止
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
