package main

import (
	"encoding/json"
	"fmt"
)

/**
 *@Author: yepeng
 *@Date: 2021/11/18 2:57 下午
 *@Name: 15switch
 *@Desc:
 */
type X struct {
	Name string `json:"name,omitempty"`
	Age  uint8  `json:"age,omitempty"`
	Temp bool   `json:"temp,omitempty"`
}

//创建一个对象

func main() {
	var x = X{
		Name: "一个X",
		Age:  12,
		Temp: false,
	}
	//序列化这个对象
	marshal, err := json.Marshal(x)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))

}
