package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now() // 返回当前时间
	p(now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC) // 使用 Date 方法构建时间
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday()) // 当前日期是是星期几？

	p(then.Before(now)) // 比较两个时间的先后
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then) // 计算两个时间差
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff)) // 一个时间可以 + 或 - 上一个时间差得到另一个时间
	p(then.Add(-diff))
}
