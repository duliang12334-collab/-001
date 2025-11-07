package main //归属包
import "fmt" //导入包

// 全剧变量
// 定义在函数外部的
var n7 = 100
var n8 = 9.7

// 一次性声明
var (
	n9  = -500
	n10 = "这是一个举例"
)

func main() {
	// 定义在「」中中的变量是局部变量
	// 第一种变量的使用方式，指定变量类型，给予赋值
	var num int = 18
	fmt.Println(num)

	//第二种，指定变量类型，但不赋值
	var num2 int
	fmt.Println(num2)

	//第三种，如果没有写变量的类型，那么根据=后面的值进行判定变量的类型（自动类型判断）
	var num3 = "tom"
	fmt.Println(num3)

	//第四种，省略var，注意 := 不能写 =
	sex := "男"
	fmt.Println(sex)

	fmt.Println("...................................")
	//支持多个变量同时声明

	//声明多个变量
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	//声明多个变量赋予多个值
	var name, n4, sex2 = "jake", 18, "男"
	fmt.Println(name, n4, sex2)

	fmt.Println(n9, n10)

}
