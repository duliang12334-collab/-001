package main

import (
	"fmt"
)

func main() {
	//定义map 变量
	var a map[int]string
	// 只声明了map，但是没有分配空间
	// 必须通过make 进行初始化
	a = make(map[int]string, 10)
	// 将键值存入
	a[001] = "张三"
	a[002] = "李四"
	fmt.Println(a)

	// 方式2
	b := make(map[int]string)
	b[201] = "aaa"
	b[202] = "bbb"
	fmt.Println(b)

	// 方式3
	c := map[int]string{
		301: "zzz",
		302: "qqq",
	}

	fmt.Println(c)

}
