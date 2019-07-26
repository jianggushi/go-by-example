package main

import "fmt"

func main() {
	var a = "initial" // 自动推导 a 为 string 类型
	fmt.Println(a)

	var b, c int = 1, 2 // 同时声明多个变量
	fmt.Println(b, c)

	var d = true // 自动推导 b 为 bool 类型
	fmt.Println(d)

	var e int // 默认零值初始化，int 类型的零值为 0
	fmt.Println(e)

	f := "apple" // 短声明方式
	fmt.Println(f)
}
