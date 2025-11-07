package main

import (
	"fmt"
)

func main() {
	// 定义一个字符串，用 len(str) 来统计字符串的字节长度
	str := "golang"
	fmt.Println(len(str))

	// new(Type) 函数进行内存分配，返回值为对应类型的指针 *int
	num := new(int)
	fmt.Printf("num类型: %T,num的值: %v,num的地址: %v,num的指针指向的值: %v", num, num, &num, *num)

}
