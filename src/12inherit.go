package main

import (
	"fmt"
)

/**
 *@Author: yepeng
 *@Date: 2021/9/9 3:38 下午
 *@Name: 12inherit
 *@Desc:go语言中模拟继承
 */

type anime struct {
	name string
}

func (a *anime) move() {
	fmt.Printf("我是动物%s 我在移动"+"\n", a.name)

}
func (a *dog) move() {
	fmt.Printf("我是小狗%s 我在移动"+"\n", a.name)
}

type dog struct {
	feet uint8
	anime
}

func (d *dog) wanwanwan() {
	fmt.Printf("我是狗狗 %v,我的feet=%v汪汪汪..."+"\n", d.name, d.feet)
}

func main() {

	var d = dog{
		feet: 1,
		anime: anime{
			name: "大黄",
		},
	}
	fmt.Println(d)
	d.wanwanwan()
	d.move()

}
