package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 统计字符串的长度，按字节进行统计：
	str := "golang你好"
	fmt.Println(len(str)) //12个字节

	fmt.Println("----------------")

	// 对字符快进行遍历：
	// 方式1: 利用for-range
	for i, value := range str {
		fmt.Printf("索引为： %d,具体值为: %c \n", i, value)

	}

	// 方式2: 利用切片 f:=[]rune(str)
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c \n", r[i])
	}

	fmt.Println("----------------")
	// 字符串转整数 strconw.Atoi
	num1, _ := strconv.Atoi("666")
	fmt.Println(num1)

	// 整数转字符串
	str1 := strconv.Itoa(88)
	fmt.Println(str1)

	// 统计一个字符串有几个指定的子串：strings.Count
	count := strings.Count("hello,world hello,golang", "hello")
	fmt.Println(count)

	// 不区分大小写进行字符串比较：strings.EqualFold
	flag := strings.EqualFold("hello", "HELLO")
	fmt.Println(flag)

	// 区分大小写进行字符串比较：==
	fmt.Println("hello" == "Hello")

	// 返回子串在字符串第一次出现的索引值，如果没有就返回-1 ： strings.index
	fmt.Println(strings.Index("hello,world hello,golang", "world"))

}
