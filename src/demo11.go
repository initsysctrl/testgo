/**
 * @Author: yp
 * @Date: 2021/7/31 7:15 下午
 * @Des:list 链表
**/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	element := l.PushBack(5)
	//l.Remove(element)
	l.InsertBefore(9999, element)
	l.Remove(element)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	var ptr *int
	fmt.Println(ptr)

}
