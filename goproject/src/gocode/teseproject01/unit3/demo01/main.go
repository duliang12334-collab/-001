package main

import "fmt"

func main() {
	// 运算符 + 的作用： 1.正数 2.数值相加 3.字符串相连
	//var n1 int = +18
	//var n2 int = 10 + 18
	//var s1 string = "ABC" + "EFG"

	//fmt.Println(n1, n2, s1)

	var n1, n2, s1 = +18, 10 + 18, "ABC" + "EFG"
	fmt.Println(n1, n2, s1)

	// 除号
	fmt.Println(10 / 3)   //两个int类型数据运算，结果一定是整数类型
	fmt.Println(10.0 / 3) //浮点类型参与运算，结果为浮点类型

	// % 取模 等价公式： a%b = a-a/b*b
	fmt.Println(10 % 3) // int整数的结果 10%3 = 10-10/3*3 = 1
	fmt.Println(10 % -3)
	fmt.Println(-10 % 3)
	fmt.Println(-10 % -3)

	// ++ 自增操作 -- 自减
	var a int = 10
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	// go语言里，++，-- 操作非常简单，只能单独使用，不能参与到运算中
	// go语言里，++，-- 只能在变量后面，不能写在变量前面 ++a , --a ❌

	//以下是关于赋值运算符
	fmt.Println("------------以下是赋值运算符示例-----------")

	// += -=
	var n3 int = 10
	n3 += 10 // n3 = n3+20
	fmt.Println(n3)
	n3 -= 5
	fmt.Println(n3)

	//交换 a b 数值
	var a1, b = 50, 100
	fmt.Printf("a1 = %v,b = %v \n", a1, b)

	//交换
	//引入一个变量
	var t int
	t = a1
	a1 = b
	b = t
	fmt.Printf("a1 = %v,b = %v \n", a1, b)

	//以下是关于关系运算符
	fmt.Println("------------以下是关系运算符示例-----------")
	fmt.Println(5 == 9) //判断左右两侧的值是否相等，返回结果：true/false
	fmt.Println(5 != 9) //不等，判断机制同上
	fmt.Println(5 < 9)
	fmt.Println(5 > 9)
	fmt.Println(5 >= 9)
	fmt.Println(5 <= 9)

	//以下是关于逻辑运算符
	fmt.Println("------------以下是逻辑运算符示例-----------")

	// 与逻辑：&& ：两个数值/表达式只要一测试false ，结果一定为false
	// 也叫短路与：只要第一个是false，后面的就不用参与运算
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)
	fmt.Println()

	// 或逻辑： ||：两个数值/表达式只要一侧是true，结果一定为 true
	//叫短路或：只要第一个是ftrue，后面的就不用参与运算
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)
	fmt.Println()

	// 非逻辑： ！ 取相反的结果
	fmt.Println(!false)
	fmt.Println(!true)

}
