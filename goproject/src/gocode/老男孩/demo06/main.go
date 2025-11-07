package main

import "fmt"

func foo(s []int) {

	s[1] = 200 //
	s = append(s, 4)
	s[0] = 100

}

func main() {
	var s = []int{1, 2, 3}
	foo(s)
	fmt.Println(s)

	// 整体执行顺序：
	//1. 先执行main函数里面的切片 【1 2 3】
	//2. 再进入foo函数内执行 s[1] = 200，因为是对同一底层数组进行改变
	// 得出结果 [1 200 3]
	//3. s = append(s, 4)对切片进行了扩容，得出了新的数组
	// 所以从这步开始，将不会对main函数做出改变，因为不是同一数组

	fmt.Println("--------------")

	(func() {
		fmt.Println("hello world")
	})()

	(func(x, y int) {
		fmt.Println(x + y)
	})(10, 10)

}
