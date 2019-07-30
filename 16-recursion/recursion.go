package main

import "fmt"

func fact(n int) int { // 递归实现 n 的阶乘
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
