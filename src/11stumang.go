package main

import "fmt"

/**
 *@Author: yepeng
 *@Date: 2021/9/6 5:17 下午
 *@Name: 11stumang
 *@Desc:学生管理系统案例
	一个系统，新增学生、更新学生、删除学生
*/
var allStudent map[int]*Student

type Student struct {
	id   int64
	name string
}

// 构造函数 创建一个学生
func NewStudent(id int64, name string) *Student {
	return &Student{id: id, name: name}
}

func main() {
	//buff := make(map[int]*Student, 48, 48)
	//buff
	fmt.Println("欢迎进入学生管理系统，请输入:")
	fmt.Println(`
1、查看所有学生
2、新增学生
3、删除学生`)
	fmt.Println("请输入你要干啥...")
	var choice = 0
	_, err := fmt.Scanln(&choice)
	if err != nil {
		return
	}
	fmt.Printf("接收到的choice=%v"+"\n", choice)
	switch choice {
	case 1:
		showAllStudent()
	case 2:
		addStudent()
	case 3:
	default:
		fmt.Println("")
	}
}

// 新增一个学生
func addStudent() {
	fmt.Print("请输入学生id")
	var id int64
	_, err := fmt.Scanln(&id)
	if err != nil {
		return
	}
	var name string
	fmt.Print("请输入学生id")
	_, err = fmt.Scanln(&name)
	if err != nil {
		return
	}
	//student := NewStudent(id, name)
	//append(allStudent,student)
}

func showAllStudent() {

}
func name() {

}
