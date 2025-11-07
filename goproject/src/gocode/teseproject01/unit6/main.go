package main

import "fmt"

func main() {
	// 利用for循环实现求和：1+2+3+4+5
	// for循环的语法格式：
	// for  初始表达式；布尔值（条件判断）；迭代因子{
	// 循环体--被执行的内容
	// }

	var sum int = 0
	// for的表达式不能用 var定义电量，要用：=
	for i := 1; i <= 5; i++ {
		sum += i
	}

	fmt.Println(sum)

	// for循环只是让提高了程序员的开发效率，底层执行效率一样

}
