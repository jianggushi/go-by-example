package main

import "fmt"

func main() {
	messages := make(chan string, 2) // 使用 make 创建带缓冲的 channel，最多可以缓冲 2 个消息

	messages <- "buffered" // 带缓冲的 channel 在 channel 满之前都不是阻塞的，可以想其中发送消息
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
