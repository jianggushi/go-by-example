package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int { // 定义 rect 类型的方法，receiver 是 指针，指针赋值修改影响外界
	return r.width * r.height
}

func (r rect) perim() int { // 定义 rect 类型的方法，receiver 是 值，值赋值修改不影响外界
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area()) // 调用者可以是值类型，调用对应的方法会转换
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area()) // 调用者可以是指针类型，调用对应的方法会转换
	fmt.Println("perim:", rp.perim())
}
