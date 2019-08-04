package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	fmt.Printf("%v\n", p) // %v 常用的默认格式

	fmt.Printf("%+v\n", p) // %+v 对于 struct 类型会带上 filed 名字

	fmt.Printf("%#v\n", p) // %#v 带源代码片段

	fmt.Printf("%T\n", p) // %T 类型

	fmt.Printf("%t\n", true) // %t bool 类型

	fmt.Printf("%d\n", 123) // %d 整数十进制形式

	fmt.Printf("%b\n", 14) // %b 整数二进制形式

	fmt.Printf("%c\n", 33) // %c 整数对应的 unicode 字符

	fmt.Printf("%x\n", 456) // %x 整数十六进制形式

	fmt.Printf("%f\n", 78.9) // %f 浮点数形式

	fmt.Printf("%e\n", 123400000.0) // %e 科学计数法形式

	fmt.Printf("%E\n", 123400000.0) // %E 科学计数法形式

	fmt.Printf("%s\n", "\"string\"") // %s 字符串形式，默认转义

	fmt.Printf("%q\n", "\"string\"") // %q 字符串形式，不转义

	fmt.Printf("%x\n", "hex this") // %x 十六进制形式

	fmt.Printf("%p\n", &p) // %p 指针形式

	fmt.Printf("|%6d|%6d|\n", 12, 345) // 占位 6 个位置，默认右对齐

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) // 占位 6 个位置，小数 2 位，默认右对齐

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) // 左对齐方式

	fmt.Printf("|%6s|%6s|\n", "foo", "b") // 占位，右对齐

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") // 占位，左对齐

	s := fmt.Sprintf("a %s", "string") // 格式化到字符串
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error") // 格式化到标准错误输出
}
