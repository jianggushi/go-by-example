package main

import "fmt"

func main() {
	var a [5]int           // 声明 a 为 [5]int 数组类型，长度也是类型的一部分
	fmt.Println("emp:", a) // 默认初始化为零值，数组元素的零值

	a[4] = 100 // 支持下标索引访问和修改数组元素的值
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a)) // 内置函数 len() 获取数组元素个数

	b := [5]int{1, 2, 3, 4, 5} // 声明时可以使用 {}，对元素进行初始化
	fmt.Println("dcl:", b)

	var twoD [2][3]int // 声明多维数组，默认初始化为元素的零值
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
