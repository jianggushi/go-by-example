package main

import "os"

func main() {
	panic("a problem") // painc 通常用来处理那些不应该发生的异常，或者不知该怎么处理的异常

	_, err := os.Create("/tmp/file") // 普通的错误，应该用 error 处理
	if err != nil {
		panic(err)
	}
}
