package main

import "fmt"

func main() {
	i := 1
	for i <= 3 { // 类似 while 循环
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ { // 经典的 for 循环
		fmt.Println(j)
	}

	for { // 类似 while true {}
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ { // 可以 continue break
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
