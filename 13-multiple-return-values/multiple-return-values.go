package main

import "fmt"

func vals() (int, int) { // 支持函数多返回值
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() // 如果想忽略某个返回值，在对应的位置设置 _
	fmt.Println(c)
}
