package main

import "fmt"

func init() {
	fmt.Println("init 1")
}
func init() {
	fmt.Println("init 2")
}
func main() {
	fmt.Println("main")
}

func init() {
	fmt.Println("init 3")
}

// output:
// init 1
// init 2
// init 3
// main
// init函数比较特殊，可以在包里被多次定义。
