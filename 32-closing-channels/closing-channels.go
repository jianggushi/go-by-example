package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs // 从 channel 中接收数据，还有第二个可选参数，代表是否关闭了 channel 且 channel 中无数据
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs") // channel 被关闭后，数据被接收完，会进入该分支
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) // 使用内置函数 close() 关闭 channel
	fmt.Println("sent all jobs")
	<-done
}
