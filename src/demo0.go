//Hello world
//包空间
package main

//导入包
import (
	"fmt"
	"strings"
)

//入口
func main() {
	s := "hello world"
	split := strings.Split(s, " ")
	fmt.Println(split[0])
	fmt.Println(split[1])
	fmt.Println("hello world")
}
