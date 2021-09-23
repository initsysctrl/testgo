package main

import (
	"fmt"
	"unicode"
)

/**
 *@Author: yepeng
 *@Date: 2021/8/23 6:01 下午
 *@Name: pointer
 *@Desc:指针
 */
func main() {
	s := "xxxx"
	fmt.Printf("value:%v type:%T address:%v \n", s, s, &s)
	fmt.Println(*&s)
	//var a *int
	//*a = 100
	//println(*a)
	i := new(int)
	*i = 1000
	println(*i)
	s2 := "deded"
	p2 := &s2
	fmt.Println(*p2)
	m := make(map[string]int, 3)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4
	s3 := "go"
	m[s3] = 999
	for k, v := range m {
		//fmt.Println(v)
		fmt.Printf("the key =%v,value=%v \n", k, v)

	}
	i2, ok := m["xxxxxx"]
	if !ok {
		fmt.Println("key is not exist")
	} else {
		fmt.Println("the value is ", i2)
	}
	//	判断一个字符串中汉子的数量
	s4 := "dasdadada大大说sdadas大大大三"
	count := 0
	for _, v := range s4 {
		fmt.Printf("char=%v"+"\n", v)
		if unicode.Is(unicode.Han, v) {
			count += 1
		}
	}
	fmt.Printf("汉字数量为%v\n", count)

}
