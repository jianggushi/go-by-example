package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select { // 增加 default 分支，使得 select 不阻塞
	case msg := <-messages: // messages 不带缓冲，阻塞接收，不成功，执行 default 分支
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg: // messages 不带缓冲，阻塞发送，不成功，执行 default 分支
		fmt.Println("send message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages: // 多分支一样的道理
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
