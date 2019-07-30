package main

import "fmt"

func intSeq() func() int { // 函数 intSeq 的返回值是一个函数类型 func() int
	i := 0
	return func() int { // 定义一个匿名函数并返回该匿名函数
		i++ // 匿名函数内部使用了外层函数的变量，形成了闭包
		return i
	}
}

func main() {
	nextInt := intSeq() // 调用 intSeq 返回一个函数类型 func() int

	fmt.Println(nextInt()) // 调用 nextInt() 返回一个 int
	fmt.Println(nextInt()) // 闭包内变量 i 会接着增长
	fmt.Println(nextInt())

	nextInts := intSeq()
	fmt.Println(nextInts())
}
