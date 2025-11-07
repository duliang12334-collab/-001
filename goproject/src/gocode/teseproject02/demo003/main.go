package main

import (
	"a2/testutils"
	"fmt"
)

var num int = test() // 全局变量

func test() int {
	fmt.Println("tese函数已被执行")
	return 10
}
func init() {
	fmt.Println("init函数已被执行") // init初始化调用
}

func main() {
	fmt.Println("main函数已被执行") //main 函数调用
	fmt.Println(testutils.Age, testutils.Name)
}
