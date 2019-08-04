package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64

	for i := 0; i < 50; i++ { // 启动 50 个 goroutine 并发执行，并发操作 ops
		go func() {
			for {
				atomic.AddUint64(&ops, 1) // 原子操作 +1
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops) // 原子操作读取 ops
	fmt.Println("ops:", opsFinal)
}
