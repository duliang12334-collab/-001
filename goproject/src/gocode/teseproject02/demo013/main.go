package main

import (
	"fmt"
)

func main() {
	//定义数组：
	var intarr [6]int = [6]int{0, 2, 4, 6, 8, 10}
	// 切片是构建在数组之上的
	// 定义一个切片名字为 slice 切出切片
	slice := intarr[1:3]

	//输出数组：
	fmt.Println("intarr:", intarr)
	//输出切片
	fmt.Println("slice:", slice)
	//输出元素个数
	fmt.Println("slice元素个数:", len(slice))
	//获取切片的容量：
	fmt.Println("slice的容量:", cap(slice))

}
