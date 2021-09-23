package main

import "fmt"

/**
 *@Author: yepeng
 *@Date: 2021/9/1 10:36 上午
 *@Name: 09struct
 *@Desc:结构体案例
 */
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func newPerson(name string, age int, gender string, hobby []string) *person {
	return &person{name: name, age: age, gender: gender, hobby: hobby}
}

func changeGender(p *person) {
	p.gender = "女"
}

func main() {
	var p = person{
		name:   "法外狂徒张三",
		age:    20,
		gender: "男",
		hobby:  []string{"足球", "篮球", "围棋"},
	}
	fmt.Println(p)
	fmt.Printf("p=%+v,type=%T"+"\n", p, p)
	changeGender(&p)
	fmt.Println(p.gender)
	p2 := new(person)
	p2.gender = "女"
	fmt.Println(p2.gender)
	p3 := &person{
		name:   "匿名",
		age:    0,
		gender: "未知",
		hobby:  nil,
	}
	fmt.Printf("%T"+"\n", p3)
	p3.name = "洛伦兹"
	fmt.Println(*p3)
	p4 := new(person)
	p4.name = "XXX"
	fmt.Printf("%+v"+"\n", p4)
	p5 := person{
		name:   "",
		age:    0,
		gender: "",
		hobby:  nil,
	}
	p5.age = 108
	fmt.Printf("%+v"+"\n", p5)
}
