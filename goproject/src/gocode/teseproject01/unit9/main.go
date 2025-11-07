package main

import "fmt"

// 自定义函数：功能：两数相加/两数相差
func cal(num1 int, num2 int) (int, int) {
	var sum int = 0
	sum = num1 + num2

	var result int = num1 - num2
	return sum, result
}

func main() {
	//调用函数：
	sum, _ := cal(100, 20)
	fmt.Println(sum)

	// 30 + 50
	//num3 := 30
	//num5 := 50
	//sum1 := cal(num3, num5)
	//fmt.Println(sum1)

}
