package main

import "fmt"

/**
 *@Author: yepeng
 *@Date: 2021/9/2 3:49 下午
 *@Name: 10funcs
 *@Desc:方法与函数的例
 */

//Dog 结构体
type Dog struct {
	name string
	age  int
}

func newDog(name string, age int) *Dog {
	return &Dog{name: name, age: age}
}
func (self Dog) wanwan() {
	fmt.Printf("我是狗，名字%v，今年%v岁，汪汪汪\n", self.name, self.age)
}
func (d Dog) Jojo() int {
	s := 12 - d.age
	if s <= 0 {
		return 0
	}
	return s
}
func (d *Dog) guonian() {
	d.age += 1
}
func (d *Dog) xxxx() {

}
func (d Dog) edede() {

}

const P = 3.1415926
const e = 2.7

func main() {
	d1 := newDog("大黄", 1)
	d2 := newDog("小白", 2)
	d1.wanwan()
	d2.wanwan()
	fmt.Printf("小狗的理论剩余寿命为 %v年"+"\n", d1.Jojo())
	d1.guonian()
	fmt.Println(*d1)
}
