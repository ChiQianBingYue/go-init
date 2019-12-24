package main

import (
	"fmt"
	"pkg/pack"
	"pkg/testutil"
)

func main() {
	fmt.Println(pack.Pack)
	fmt.Println(testutil.Util)
}

// init test_util
// init pack  5
// 6
// 5
// 由于pack包的初始化依赖test_util，因此运行时先初始化test_util再初始化pack包；

// 如果　pack 不依赖　testutil，　则打印
// init pack
// init test_util
// 6
// 5
