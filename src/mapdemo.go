package main

import (
	"fmt"
	"strings"
)

/**
 *@Author: yepeng
 *@Date: 2021/8/23 6:21 下午
 *@Name: mapdemo
 *@Desc:map 练习题,统计一个字符串中每个字符出现的次数
 */
func main() {
	//m := make(map[string]int, 10) //预估容量
	//m["a"] = 100
	//m["b"] = 99
	//m["c"] = 88
	//fmt.Println(m)
	//for k, v := range m {
	//	fmt.Println(k, ":", v)
	//}
	//delete(m, "a")
	//fmt.Println(m)
	s := "hello world,hello golang"
	m := make(map[string]int, 10)
	for _, v := range s {
		//fmt.Printf("%c\n", v)
		if v == ' ' {
			continue
		}
		cs := string(v)
		m[cs] += 1

	}
	fmt.Println(m)
	fmt.Println(strings.Count("hello world", "o"))
	m2 := make(map[int]int)
	for _, v := range m2 {
		fmt.Printf("v=%v"+"\n", v)
	}
}
