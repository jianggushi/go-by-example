package main

import "fmt"

func main() {
	m := make(map[string]int) // m 为 map[string]int 的 map 类型，类似 python 中的 dict 及其他语言中的 hash table

	m["k1"] = 7 // 可以使用 name[key] = val 的方式添加或更新 key-value
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"] // 可以使用 name[key] 访问 map 中的元素，如果 key 不存在，返回值类型的默认零值
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m)) // 内置函数 len() 返回 map 中 key-value 元素对个数

	delete(m, "k2") // 内置函数 delete() 删除指定 key 元素
	fmt.Println("map:", m)

	_, prs := m["k2"] // 访问 map 中的元素 name[key]，还会返回可选的第二个参数，bool 类型，代表 key 是否存在
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2} // 也可以使用字面值直接初始化
	fmt.Println("map:", n)
}
