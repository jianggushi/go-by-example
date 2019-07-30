package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine") // goroutine 协程是一种轻量级的用户级线程，可以使用关键字 go 来启动

	go func(msg string) { // 函数也可以是匿名函数
		fmt.Println(msg)
	}("going")

	fmt.Scanln() // 阻塞等待用户输入
	fmt.Println("done")
}
