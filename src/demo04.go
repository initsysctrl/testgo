package main

import (
	"fmt"
	"strings"
)

/**
 *@Author: yepeng
 *@Date: 2021/8/22 7:41 下午
 *@Type: demo4
 *@Desc:字符串处理
 */

func main() {
	s := "hello world"
	split := strings.Split(s, " ")
	fmt.Println(split[0])
	println(split[1])
}
