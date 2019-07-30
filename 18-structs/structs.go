package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20}) // 按顺序初始化赋值，不能省略或间隔

	fmt.Println(person{name: "Alice", age: 30}) // 指定元素名初始化

	fmt.Println(person{name: "Fred"}) // 未初始化的元素默认初始化为元素类型零值

	fmt.Println(&person{name: "Ann", age: 40}) // 操作符 & 得到 struct 指针

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // 通过操作符 . 访问 struct 中的元素

	sp := &s            // sp 为 struct 指针
	fmt.Println(sp.age) // 指针类型同样通过操作符 . 访问元素

	sp.age = 51 // 可以修改元素的值
	fmt.Println(sp.age)
}
