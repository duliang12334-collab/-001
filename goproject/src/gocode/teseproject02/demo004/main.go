package main

import "fmt"

// 函数功能：求和
// 函数的名字： getSunm 参数为空 返回值为函数func 函数为int类型，返回值也是int类型
func getSum() func(int) int {
	var sum int = 10
	return func(num int) int {
		sum = sum + num
		return sum
	}
}

// 闭包：返回匿名函数+匿名函数以外的变量num
func main() {
	s := getSum()
	fmt.Println(s(10))
	fmt.Println(s(10))
	fmt.Println(s(10))

}
