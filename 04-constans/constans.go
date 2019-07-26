package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000 // 无类型常量

	const d = 3e20 / n // 无类型常量可以表示更大精度，不少于 128bit
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
