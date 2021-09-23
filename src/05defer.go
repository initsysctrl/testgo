package main

import "fmt"

/**
 *@Author: yepeng
 *@Date: 2021/8/30 1:03 上午
 *@Name: 05defer
 *@Desc:
 */

func main() {
	fmt.Println("start...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end...")
}
