/**
 * @Author: yp
 * @Date: 2021/7/30 2:56 下午
 * @Des:结构体的方法，值接收和指针接收
**/
package main

import "fmt"

type User struct {
	name string
	age  int
}

//接收对象是一个值对象，调用方法后是拷贝，但不修改
func (user User) setName1(name string) {
	user.name = name

}

//接收对象是指针，调用方法后也是拷贝，但可以修改
func (user *User) setName2(name string) {
	user.name = name

}

func main() {
	user := User{
		name: "jojo",
		age:  18,
	}
	user.setName1("yp") //这里没有改变user的值
	fmt.Println(user)
	user.setName2("yp") //这里改变了
	fmt.Println(user)
}
