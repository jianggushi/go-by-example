package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // 关闭非空 channel 后，channel 中还未被接收的数据仍然保持可以被接收

	for elem := range queue { // 使用 for-range channel 当 channel 被关闭并且没有数据，自动结束 for-range
		fmt.Println(elem)
	}
}
