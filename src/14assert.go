package main

import "fmt"

/**
 *@Author: yepeng
 *@Date: 2021/9/24 2:52 下午
 *@Name: 14assert
 *@Desc:
 */
func assert(a interface{}) {
	switch a.(type) {
	case string:
		fmt.Println("a type is string")
	case int, int32, int8, int16, int64:
		fmt.Println("a type is int")
	case float32, float64:
		fmt.Println("a type is float")
	case bool:
		fmt.Println("a type is bool")
	default:
		fmt.Println("unknown type..")
	}

}
func main() {
	var v = 0x1
	assert(v)
}
