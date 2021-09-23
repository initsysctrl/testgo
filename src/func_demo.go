package main

import "fmt"

/**
 *@Author: yepeng
 *@Date: 2021/8/23 7:39 下午
 *@Name: func_demo
 *@Desc:
 */
func f1(x int, y int) (ret int) {
	ret = x + y
	return ret
}
func main() {
	ret := f1(123, 1321)
	fmt.Println(ret)

}
