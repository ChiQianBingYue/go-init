package main

import "fmt"

func init() {
	fmt.Println("init")
}
func main() {
	init()
}

// output:
// # callinit
// ./main.go:9:2: undefined: init
