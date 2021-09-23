package main

import "fmt"

/**
 *@Author: yepeng
 *@Date: 2021/8/23 5:27 下午
 *@Name: make_slice
 *@Desc:切片
 */
func main() {
	s := make([]string, 0, 4)
	s = append(s, "sb")
	fmt.Println(s)
	arr := []int{1, 2, 3}
	arr = append(arr, 9999)
	fmt.Println(arr)
}
