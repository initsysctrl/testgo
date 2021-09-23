//变量与常量
package main

import (
	"fmt"
)

//标准命名
var name1 string = "jojo"

//初始化后可以省略类型
var name2 = "jon"

// PI 常量
const PI = 3.1415926
const (
	p1 = 2312312
	p2 = 312312312
)
const (
	a1 = iota //a1=0
	a2 = 100  //a2=100
	_         //废弃
	a3        // a3=2
)
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func test1() {
	//方法里面的变量可以用海象运算符 省略关键字var
	a := 1
	var b int
	println(a, b)
}

func main() {
	fmt.Println(KB, MB, GB, TB, PB)
	i := 077

	fmt.Printf("%T\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%b\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%x\n", i)

}
