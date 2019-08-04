package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs) // 原地排序 []string
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints) // 原地排序 []int
	fmt.Println("Ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
