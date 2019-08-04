package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5) // 构造 5 个 request
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter // 通过 limiter 限制每隔 200ms 处理一个 request
		fmt.Println("request", req, time.Now())
	}

	burstyRequests := make(chan int, 5) // 同样构造 5 个 request
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	burstyLimiter := make(chan time.Time, 3) // limiter 待缓冲，长度 3
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	for req := range burstyRequests { // 处理速度也会加快
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
