package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print(rand.Intn(100), ",") // 返回 [0,100) 内的随机数，没有设置种子，每次运行返回的数字相同
	fmt.Println(rand.Intn(100))

	fmt.Println(rand.Float64()) // 返回 [0,1] 内的随机数

	fmt.Print((rand.Float64()*5)+5, ",") // 返回 [5,10] 内的随机数
	fmt.Println((rand.Float64() * 5) + 5)

	s1 := rand.NewSource(time.Now().UnixNano()) // 用时间戳当做种子
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",") // 每次运行返回的数字随机
	fmt.Println(r1.Intn(100))

	s2 := rand.NewSource(42) // 用固定数字当做种子
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",") // 每次运行返回的数字固定
	fmt.Println(r2.Intn(100))

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Println(r3.Intn(100))
}
