package main

//import "fmt"
//import "unsafe"
//一次倒入多个包：
import (
	"fmt"
	"unsafe"
)

func main() {
	//定义一个整数类型
	var num1 int8 = 23
	fmt.Println(num1)

	// printf函数：打印类型
	var num2 = 28
	fmt.Printf("num2类型是:%T", num2)
	fmt.Println()

	//占用了多少字节
	fmt.Println(unsafe.Sizeof(num2))

	//选择合适的数据类型，节省内容空间
	//例如：var age int64 = 18,占用字节太大，浪费内存，可取代为：uint8 或byte
	var age uint8 = 18
	fmt.Println(age)

}
