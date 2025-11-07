package main

import "fmt"

func main() {
	// 实现功能：如果余额小于1，提示：余额不足
	//var count int = 0
	//单分支：
	//if count < 1 {
	//	fmt.Println("对不起,余额不足")
	//}

	// if后面表达式，返回结果一定是true或者false
	// 如果返回结果是true的话，那么{}中的代码会执行
	// 如果返回是false的话，那么{}中的代码就不会执行
	// if后面一定要有空格，和条件表达式分割开来
	// {}一定不能省略

	// 在golong里，if后面可以并列的加入变量的定义：

	//if count := 10; count < 1 {
	//	fmt.Println("对不起,余额不足")
	//} else {
	//	fmt.Println("下单成功")
	//}

	// 多分支

	var amount int = 8189999
	if amount >= 10000 {
		fmt.Println("钻石用户")
	} else if amount >= 1000 {
		fmt.Println("黄金用户")
	} else if amount >= 100 {
		fmt.Println("白银用户")
	} else if amount >= 10 {
		fmt.Println("黄铜用户")
	} else {
		fmt.Println("请充值")
	} // 建议保存最后的else，形成闭环

}
