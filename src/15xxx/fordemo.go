package main

import "fmt"

/**
 *@Author: yepeng
 *@Date: 2021/11/18 3:04 下午
 *@Name: fordemo
 *@Desc:
 */
func main() {
	for a := 0; a < 5; a++ {
		fmt.Println(a)
	}
	arr := make([]string, 0)
	arr = append(arr, "a")
	arr = append(arr, "b")
	arr = append(arr, "c")
	arr = append(arr, "d")
	arr = append(arr, "e")
	arr = append(arr, "f")
	arr = append(arr, "g")

	for index, value := range arr {
		fmt.Println(index, value)
	}

	m := make(map[int]string)
	m[1] = "xxx"
	m[2] = "yyy"
	m[3] = "zzz"

	for key, value := range m {
		fmt.Println(key, value)
	}
}
