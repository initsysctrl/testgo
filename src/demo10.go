/**
 * @Author: yp
 * @Date: 2021/7/31 6:53 下午
 * @Des:map
**/
package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[string]int{}
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	for k, v := range m {
		fmt.Printf("%s -> %d\n", k, v)
	}
	arr := []int{1, 2, 3}
	for i, v := range arr {
		fmt.Printf("index %d -> value %d\n", i, v)
	}
	delete(m, "a")
	for k, v := range m {
		fmt.Printf("%s -> %d\n", k, v)
	}
	//	安全的syn.map
	m2 := sync.Map{}
	m2.Store("ccc", 1)
	m2.Store("ddd", 2)
	m2.Store("eee", 3)
	fmt.Println(m2.Load("ccc"))
	fmt.Println(m2.Load("fff"))
	m2.Delete("eee")
	//Range() 方法可以遍历 sync.Map，遍历需要提供一个匿名函数，参数为 k、v，类型为 interface{}，每次 Range() 在遍历一个元素时，都会调用这个匿名函数把结果返回。
	m2.Range(func(key, value interface{}) bool {
		fmt.Println("key=", key, " value=", value)
		return true
	})
}
