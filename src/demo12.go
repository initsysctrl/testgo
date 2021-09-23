/**
 * @Author: yp
 * @Date: 2021/7/31 7:35 下午
 * @Des:流程控制
**/
package main

import "fmt"

func main() {
	v := 2
	switch v {
	case 1:
		break
	case 2:
		break
	case 3:
		break
	default:
		break

	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		for j := 0; j < 10; j++ {
			if j == 5 {
				goto breakFlag
			}
		}
	}
breakFlag:
	fmt.Println("break break!")
}
