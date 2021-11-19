package main

import (
	"encoding/json"
	"fmt"
)

/**
 *@Author: yepeng
 *@Date: 2021/11/18 3:30 下午
 *@Name: funcdede
 *@Desc:
 */

type Dog struct {
	Name string `json:"name"`
	sex  string `json:"sex"`
	Pig  int    `json:"pig"`
}

func NewDog(name string, sex string, pig int) *Dog {
	return &Dog{Name: name, sex: sex, Pig: pig}
}

func main() {
	var dog = Dog{
		Name: "大黄",
		sex:  "雄性",
	}
	fmt.Println(dog)
	marshal, err := json.Marshal(dog)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
	//dog.Action()
}
