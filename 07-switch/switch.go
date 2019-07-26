package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i { // 经典的 switch case 语句，case 中不需要 break
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // case 语句中支持多个条件
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's the weekday")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) { // type switch 对 interface 的真实类型做判断
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
