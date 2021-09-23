/**
 * @Author: yp
 * @Date: 2021/7/31 5:47 下午
 * @Des:数组
**/
package main

import "fmt"

func main() {
	arr := []int{1, 2, 34, 4}
	arr1 := [...]int{1, 2, 34, 4}
	arr2 := []int{0: 888, 1: 9999}

	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)
	//	数组的遍历
	for i, v := range arr {
		fmt.Println(i, ":", v)
	}
	arr[0] = 9999
	fmt.Println(arr)
	b := arr
	b[0] = 8888
	fmt.Println(arr)
	fmt.Println(b)

	arr = []int{1, 2, 3, 4}
	sum := 0
	for _, value := range arr {
		sum += value
	}
	println(sum)

}
