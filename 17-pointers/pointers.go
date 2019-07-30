package main

import "fmt"

func zeroval(ival int) { // 参数类型是 int 类型，传参是传值
	ival = 0 // 在函数内部改变参数的值，对外界无影响
}

func zeroptr(iprt *int) { // 参数类型是 *int 类型，指针，传参是参数地址
	*iprt = 0 // 在函数内部改变指针的值，会影响外界
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroprt:", i)

	fmt.Println("pointer:", &i)
}
