package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串替换 / strings.Replace -1 是更换全部
	str1 := strings.Replace("hello,world hello,golang hello,hello", "hello", "Hi", -1) // -1 更换全部
	str2 := strings.Replace("hello,world hello,golang hello,hello", "hello", "Hi", 2)
	fmt.Println(str1)
	// 输出：Hi,world Hi,golang Hi,Hi
	fmt.Println(str2)
	// 输出：Hi,world Hi,golang hello,hello

	// 按照指定的某个字符（自定义字符）为分割标识，切割为字符串数据： dtrings.Split
	arr := strings.Split("go-python-jave", "-")
	fmt.Println(arr)
	// 输出 [go python jave]

	// 将字符串进行大小写转换： ToLower / ToUpper
	fmt.Println(strings.ToLower("Go"))
	fmt.Println(strings.ToUpper("go"))

}
