package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { // range 遍历 array 或者 slice，返回元素的 index 和 value
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"} // range 遍历 map，返回元素的 key 和 value，返回元素的顺序是不定的
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs { // range 遍历 map，可以省略返回的元素的 value
		fmt.Println("key:", k)
	}

	for i, c := range "go 世界" { // range 遍历 string，返回 string 中的每个字符的索引和 unicode 码点
		fmt.Println(i, c)
	}
}
