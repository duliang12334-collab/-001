package main

import (
	"fmt"
)

func main() {
	// 登陆
	var username, password string
	fmt.Println("请输入用户名和密码")
	fmt.Scan(&username, &password)
	if username == "luozhengyi123" && password == "qwe123.." {
		// 分数
		var score int
		fmt.Println("请输入你的成绩：")
		fmt.Scan(&score)

		if score < 0 || score > 100 {
			fmt.Println("你输入的成绩有误")
		} else if score > 90 {
			fmt.Println("成绩优秀")
		} else if score > 60 {
			fmt.Println("成绩合格")
		} else {
			fmt.Println("成绩不合格")
		}

		fmt.Println("show score end")

	} else {
		fmt.Println("用户名或密码错误")
	}

}
