package main

import (
	"fmt"
)

//经典代码，打印九九乘法表
func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j > i {
				fmt.Print("\n")
				break
			}
			fmt.Print(i, "*", j, "=", i*j, " ")
		}
	}
}
