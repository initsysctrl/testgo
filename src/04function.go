package main

import (
	"fmt"
	"time"
)

/**
 *@Author: yepeng
 *@Date: 2021/8/30 12:17 上午
 *@Name: 04function
 *@Desc:
 */

func time_spend(f func()) int64 {
	start := time.Now().UnixNano() / 1e6
	f()
	end := time.Now().UnixNano() / 1e6
	fmt.Println(end - start)
	fmt.Printf("start=%v end=%v spent=%v"+"\n", start, end, end-start)
	return end - start
}
func ttt() {
	sum := 0
	for i := 0; i < 10000000; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func main() {
	var f = func(x int, y int) int {
		return x * y
	}
	f(12, 123)
	time_spend(ttt)
}
