package main

import "fmt"

func main() {
	if 7%2 == 0 { // if 中的条件表达式必须是 bool 类型的
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 { // 可以只有 if 语句，省略 else 语句
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 { // 条件表达式之前可以声明变量，变量可以在 if else 语句中使用
		fmt.Println("num is negative")
	} else if num < 10 {
		fmt.Println("has 1 digit")
	} else {
		fmt.Println("has multiple digits")
	}
}
