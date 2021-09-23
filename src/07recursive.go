package main

import "fmt"

/**
 *@Author: yepeng
 *@Date: 2021/8/30 3:11 下午
 *@Name: 07recursive
 *@Desc:递归函数,经典斐波那契数列
 */
var m = make(map[int]int, 81)

func recur(i int) int {
	switch i {
	case 0:
		m[i] = 0
		return 0
	case 1:
		m[i] = 1
		return 1
	case 2:
		m[i] = 2
		return 2
	default:
		return recur(i-1) + recur(i-2)
	}
}
func main() {
	for i := 0; i < 10; i++ {
		x := recur(i)
		fmt.Println(x)
	}
}
