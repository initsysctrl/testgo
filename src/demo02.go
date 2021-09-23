package main

//一个判断数字正负的函数
func test(a int) {
	if a > 0 {
		println("a >0")
	} else if a < 0 {
		println("a<=")
	} else {
		println("a==0")
	}
}
func main() {
	test(1)

}
