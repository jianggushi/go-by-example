package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ { // 启动 3 个 worker，并发执行 job
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ { // 生成 5 个 job
		jobs <- j
	}
	close(jobs) // 生成 5 个 job 后 关闭 channel，让 worker 能够结束

	for a := 1; a <= 5; a++ {
		<-results
	}

}
