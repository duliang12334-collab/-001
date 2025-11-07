package main

import (
	"fmt"
)

// 关于指针
func main() {
	var age int = 18
	// &符号+变量 就可以获取这个变量的地址
	fmt.Println(&age) //0xc00009c008

	//定义一个指针的变量
	//var 代表声明一个变量
	//ptr 是指针变量的名字
	//ptr 对应的类型： *int是一个指针类型（可以理解为指向 int类型的指针）
	//&age 就是一个地址：是 prt变量的具体的值
	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr内存地址:", &ptr)

	//想获取ptr这个指针或者这个地址指向的那个数据：
	fmt.Printf("ptr的数值是: %v \n", *ptr)

	// 可以通过指针改变指向的值
	var z int = 100
	fmt.Println(z)

	var prt01 *int = &z
	*prt01 = 50
	fmt.Println(z)

}
