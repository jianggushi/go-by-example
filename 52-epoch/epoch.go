package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	secs := now.Unix()        // 获取 unix 时间戳，单位 s
	nanos := now.UnixNano()   // 单位 ns
	millis := nanos / 1000000 // 单位 ms
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0)) // 根据 unix 时间戳获取时间格式
	fmt.Println(time.Unix(0, nanos))
}
