package main

import (
	"fmt"
	"time"
)

func main() {

	// 时间和日期的函数，需要倒入time包，调用函数 time.Now
	now := time.Now()
	// Now() 返回值是一个结构体 类型为：time.Time
	fmt.Printf("%v 对应的类型为： %T \n", now, now)

	// 调用不同的恶类型用不同的 函数
	// now.Year / now.Month / Day / Hour / Minute / Second 等等

	fmt.Println("-----------------------")

	// 利用 Printf将 日期以字符串的形式进行输出：
	fmt.Printf("日期： %d-%d-%d 时间：%v:%v:%v \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	// 利用 Sprintf 对字符串赋值，反复调用
	datestr := fmt.Sprintf("日期： %d-%d-%d 时间：%v:%v:%v \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(datestr)

	// 更简便的方法 now.Format ("2006/01/02 15:04:05") 固定格式！！！！
	datestr2 := now.Format("2006/01/02 15:04:05")
	fmt.Println(datestr2)

}
