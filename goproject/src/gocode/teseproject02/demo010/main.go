package main

import (
	"fmt"
)

func main() {
	test()
	fmt.Println("如上述流程成功则执行下一步")
	fmt.Println("正常执行")

}

func test() {

	// 利用 defer+recover 来捕获错误进行打印后继续进行下一步操作
	defer func() { //加匿名函数
		// 调用recover 内置函数，进行错误捕获
		err := recover() //如果没有错误，则返回值为0 /nil
		if err != nil {  //如果返回不是nil，不等于0 则打印具体错误信息
			fmt.Println("错误已被捕获 err是: ", err)
		}
	}() // 加() 确保函数被调用
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}
