package main

import "fmt"

func main() {
	fmt.Println("go" + "lang") // 字符串可以通过 + 进行连接

	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("7.0 / 3.0 =", 7.0/3.0)

	fmt.Println(true && false) // false
	fmt.Println(true || false) // true
	fmt.Println(!true)
}
