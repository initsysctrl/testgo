package main

import "fmt"

/**
 *@Author: yepeng
 *@Date: 2021/8/31 11:53 上午
 *@Name: 08
 *@Desc:走台阶问题的经典实现
 */
//核心方法，迭代递归
func core(i int, m map[int]int) (int, map[int]int) {
	//增加一个buff来保存值
	if m[i] != 0 {
		return m[i], m
	}
	if i == 0 {
		m[0] = 0
		return 0, m
	} else if i == 1 {
		m[1] = 1
		return 1, m
	} else if i == 2 {
		m[2] = 2
		return 2, m
	} else {
		ret1, _ := core(i-1, m)
		ret2, _ := core(i-2, m)
		m[i] = ret1 + ret2
		return ret1 + ret2, m
	}
}

func stair(n int) (int, map[int]int) {
	//构建一个map，来保存相关的值，避免重复计算
	m := make(map[int]int, n)
	return core(n, m)
}
func main() {
	ret, buff := stair(10)
	fmt.Println(ret)
	fmt.Println(buff)
}
