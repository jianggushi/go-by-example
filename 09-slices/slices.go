package main

import "fmt"

func main() {
	s := make([]string, 3) // s 为 []string 切片类型，切片类型和数组很像，但是 [] 中不需要指定长度
	fmt.Println("emp:", s) // 可以使用内置函数 make() 声明和初始化切片类型，初始化为元素零值

	s[0] = "a" // 和数组类型一样，slice 类型可以通过下标索引访问和修改元素的值
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s)) // 内置函数 len() 获取切片中元素个数

	s = append(s, "d") // slice 类型是可变长的，可以通过内置函数 append() 扩展
	s = append(s, "e", "f", "g")
	fmt.Println("apd:", s)

	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s)) // 内置函数 cap() 获取切片的容量，容量 >= 长度

	c := make([]string, len(s))
	copy(c, s) // 内置函数 copy() 可以对 slice 进行复制
	fmt.Println("cpy:", c)

	l := s[2:5] // 可以对 slice 进行切片操作
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"} // 也可以通过 {} 的方式声明初始化 slice 类型变量
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3) // 多维 slice 使用 make() 初始化时，指定的是外层 slice 的长度
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen) // 内层 slice 需要单独初始化，每个 slice 的长度可以不一样
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
