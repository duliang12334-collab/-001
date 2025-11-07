package main

import (
	"fmt"
)

func main() {
	//定义数组
	var a1 [6]int = [6]int{1, 2, 3, 4, 5, 6}
	// 定义切片：
	var slice []int = a1[1:4]
	fmt.Println(slice)
	fmt.Println(a1)

	slice2 := append(slice, 666, 999)
	fmt.Println(slice2)

	slice = append(slice, 666, 999)
	fmt.Println(a1)

	fmt.Println("--------------------")
	var b1 []int = make([]int, 10)
	copy(b1, slice)
	fmt.Println(b1)

}
