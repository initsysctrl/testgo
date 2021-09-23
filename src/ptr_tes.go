package main

import (
	"flag"
	"fmt"
)

//const PH = 0.7

var mode = flag.String("mode", "default model", "run model")

func main() {
	flag.Parse()
	fmt.Println(*mode)
	p := new(string)
	*p = "what's the fuck"
	print(*p)
}
