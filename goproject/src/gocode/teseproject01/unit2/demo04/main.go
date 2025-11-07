package main

import "fmt"

func main() {
	//定义浮点类型
	var num1 float32 = 3.14
	fmt.Println(num1)
	//负数
	var num2 float32 = -3.14
	fmt.Println(num2)

	// float64/科学计数法 E/e 都行
	var num3 float32 = 3.14e+2
	fmt.Println(num3)

	var num4 float32 = 3.14e-2
	fmt.Println(num4)

	//浮点数会存在精度损失，建议使用float64
	//golang中默认的浮点使用类型为：float64

}
