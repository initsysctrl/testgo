/**
 * @Author: yp
 * @Date: 2021/8/1 7:14 下午
 * @Des:接口
**/
package main

import "fmt"

type Invoker interface {
	// Call 需要实现一个call方法
	Call(interface{})
}

// Struct 定义一个结构体
type Struct struct {
}

// Call 实现接口方法
func (s *Struct) Call(i interface{}) {
	//panic("implement me")
	fmt.Println("from struct", i)
}

type FunCaller func(interface{})

func main() {
	var invoker Invoker = new(Struct)
	invoker.Call("hello")
}
