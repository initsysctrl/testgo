//for 循环
package main

import "fmt"

func main() {
	//经典的for循环
	for i := 0; i < 10; i++ {
		println("i=", i)
	}
	//带range的循环，遍历一个数组
	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Printf("index=%d value=%d \n", i, v)
	}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	println(sum)
	//死循环
	for {
		println("pa pa pa")
		break
	}

}
