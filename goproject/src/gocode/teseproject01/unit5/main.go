package main

import "fmt"

func main() {
	//实现功能，根据累计充值金额，判断会员登记
	// >= 10000 钻石会员
	// >= 1000 黄金会员
	// >= 100 白银会员
	// >= 10 青铜会员
	// >= 0 未开通会员

	// 给出一个累积充值金额
	var amount int = 96
	// 根据金额判断等级：
	// switch 后面是一个表达式，表达式的结果一次跟case进行比较，满足结果的话就执行冒号后面的代码
	// default 是用来“兜底”的一个分支，其他case分支都不走的情况下会走default分支
	// default 分支可以放在任意位置上，不一定非要放在最后
	switch amount / 10 {
	case 10:
		fmt.Println("钻石会员")
	case 9:
		fmt.Println("黄金会员")
	case 8:
		fmt.Println("白银会员")
	default:
		fmt.Println("请充值")

	}

}
