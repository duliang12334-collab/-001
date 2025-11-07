package main

import "fmt"

func Add(num1 int, num2 int) int {
	defer fmt.Println("num1=", num1) // defer推迟执行，已栈点形式先进后出
	defer fmt.Println("num2=", num2)

	var sum int = num1 + num2
	fmt.Println("sum=", sum)
	return sum
}

func main() {
	fmt.Println(Add(40, 60))
}
