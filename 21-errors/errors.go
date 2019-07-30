package main

// see more: https://blog.golang.org/error-handling-and-go

// error 类型是一个 built-in 的 interface 类型
// type error interface {
// 	Error() string
// }

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) { // 函数支持多返回值，通过一个 error 返回值，来明确的传达函数是否返回了错误
	if arg == 42 {
		return -1, errors.New("can't work with 42") // 通过 errors.New() 方法可以定义一个 error
	}
	return arg + 3, nil // 如果没有错误参数，返回 nil，nil 代表没有错误
}

type argError struct { // 自定义的类型 argError 只要实现了 Error() 方法就可以作为 error 接口类型
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil { // in-line check 方式，检查 error
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok { // f2 返回的是 error 接口类型，获取真实类型需要进行类型断言
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
