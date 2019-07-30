package main

import "fmt"

func sum(nums ...int) { // 支持函数可变参数，可变参数组成一个 slice
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...) // 如果参数原本是 slice，需要使用 slice... 的方式传参
}
