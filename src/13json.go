package main

import (
	"encoding/json"
	"fmt"
)

/**
 *@Author: yepeng
 *@Date: 2021/9/9 4:13 下午
 *@Name: 13json
 *@Desc:序列化和反序列化
 */

type X struct {
	Name string `json:"name,omitempty"`
	Age  uint8  `json:"age,omitempty"`
	Temp bool   `json:"temp,omitempty"`
}

func main() {
	var x = X{
		Name: "一个X",
		Age:  12,
		Temp: false,
	}
	//序列化
	marshal, err := json.Marshal(x)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
	// 反序列化
	s := `{"name":"这是大叔","age":77}`
	var x2 X
	err = json.Unmarshal([]byte(s), &x2)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(x2)
	// 再次序列化
	bytes, _ := json.Marshal(x2)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
}
