package main

import (
	"fmt"
	"sort"
)

type byLength []string // 只要实现 sort.Sort 这个 interface 类型的几个接口 Len(), Swap(), Less()

func (s byLength) Len() int           { return len(s) }
func (s byLength) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byLength) Less(i, j int) bool { return len(s[i]) < len(s[j]) }

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
