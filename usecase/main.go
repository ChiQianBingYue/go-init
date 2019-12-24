package main

import (
	"fmt"
	_ "net/http/pprof"
)

var initArg [20]int

func init() {
	initArg[0] = 10
	for i := 1; i < len(initArg); i++ {
		initArg[i] = initArg[i-1] * 2
	}
}

func main() {
	fmt.Println(initArg)
}

// init函数的主要用途
// 1. 初始化不能被初始化表达式初始化的变量,　如数组的值
// 2. import _ "net/http/pprof"   golang对没有使用的导入包会编译报错，但是有时我们只想调用该包的init函数，不使用包导出的变量或者方法，这时就采用上面的导入方案。
