package main

import "fmt"

func main() {
	//实现功能：键盘录入学生的姓名，年龄，成绩，是否为VIP
	//方式1 ：Scanln

	var name string
	//fmt.Println("姓名:  ")
	//fmt.Scanln(&name) // &name:传入地址的目的是为了地址中的值改变后影响外面变量的值

	var age int
	//fmt.Println("年龄:  ")
	//fmt.Scanln(&age)

	var score float32
	//fmt.Println("成绩:  ")
	//fmt.Scanln(&score)

	var isVIP bool
	//fmt.Println("是否为VIP:  ")
	//fmt.Scanln(&isVIP)

	// 在控制台打印上述数据：
	//fmt.Printf("姓名:%v, 年龄:%v,成绩: %v,VIP:%v", name, age, score, isVIP)

	// 方法2: Scanf
	fmt.Println("姓名， 年龄， 成绩， 是否是VIP, ")
	fmt.Scanf("%s %d %f %t", &name, &age, &score, &isVIP)
	fmt.Printf("姓名:%v, 年龄:%v,成绩: %v,VIP:%v", name, age, score, isVIP)

}
