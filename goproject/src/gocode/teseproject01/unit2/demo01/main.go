package main

import "fmt"

func main() {
	//1.变量的声明
	var age int
	//2.变量的赋值
	age = 28
	//变量的使用
	fmt.Println("age = ", age)

	//声明和赋值可以合成一句
	var age2 int = 38
	fmt.Println("age2 = ", age2)

	//不能重复声明
	//var age2 int = 48
	//fmt Println("age2 = ", age2)

	//不能使用未赋值的变量
	var age3 int
	fmt.Println("age3 = ", age3)
}
