package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int { // 连续多个类型相同的参数，可以用逗号间隔，只写一次类型
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusplus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)
}
