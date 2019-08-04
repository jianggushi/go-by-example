package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach") // MatchString 正则匹配字符串
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch") // 先编译正则表达式
	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch")) // 找到第一个匹配的返回

	fmt.Println(r.FindStringIndex("peach punch")) // 找到第一个匹配的首尾位置

	fmt.Println(r.FindStringSubmatch("peach punch")) // 返回包含 () 中的子匹配

	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // 返回包含 () 中的子匹配的首尾位置

	fmt.Println(r.FindAllString("peach punch pinch", -1)) // 返回所有匹配

	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach"))) // 匹配 []byte

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // 匹配并替换

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper) // 自定义替换 function
	fmt.Println(string(out))
}
