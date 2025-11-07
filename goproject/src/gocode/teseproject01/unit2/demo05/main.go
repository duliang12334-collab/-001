package main

import "fmt"

func main() {
	//定义字符类：
	var c1 byte = 'a'
	fmt.Println(c1) //97
	var c2 byte = '6'
	fmt.Println(c2) //54
	var c3 byte = '<'
	fmt.Println(c3) //60

	// 字符类型，本质上就是一个整数，可以直接参与运算，输出字符的时候，会将对应的马值做一个输出
	//字母，数字，标点等字符，底层是按照ASCII进行储存的
	// golang的字符对应使用的是UTF-8编码
	//如果打印中文的“中”字，码值为20013，byte类型益处范围，可以用int类型。例如：
	// var c4 byte = '中‘ ❌ 报错，正确如霞：
	var c4 int = '中'
	//打印时选择类型,选择格式化进行输出
	fmt.Printf("c4是：%c\n", c4)

	fmt.Println("..........下列是关于转义字符.............")
	//练习转义字符
	// \n 换行
	fmt.Println("aaaa\nbbbb")
	// \b 退格
	fmt.Println("aaaa\bbbbb")
	// \r 光标回到本行开头，后续输入就会替换原有字符
	fmt.Println("aaaa\rbbbb")
	// \t 制表符
	fmt.Println("aaaabbbb")
	fmt.Println("aaaa\tbbbb")
	fmt.Println("aaaaaaaa\tbbbb")
	// \" 标注“golang”
	fmt.Println("\"golang\"")

	fmt.Println("---------下列是关于bool值----------")

	// 测试布尔值类型的数值
	var F1 bool = true
	fmt.Println(F1)

	var F2 bool = false
	fmt.Println(F2)

	var F3 bool = 5 < 9
	fmt.Println(F3)

	var F4 bool = 5 > 9
	fmt.Println(F4)

	fmt.Println("---------下列是关于字符串----------")
	//1.定义一个字符串
	var s1 string = "你好，全面拥抱GOLANG"
	fmt.Println(s1)

	//2.字符串是不可变的：指的是字符串一旦定义好，其中的字符值不能改变
	var s2 string = "abc"
	//s2 = "def" 可以打印
	//s2[0] = 't' 无法改变
	fmt.Println(s2)

	//3.字符串的表示形式：
	//如果字符串中没有特殊字符，字符串的表示形式用双引号
	//var s3 string = "abcdefg"
	//如果字符串中有特殊字符，字符串的表示形式用反引号``
	var s4 string = `
	/ var c4 byte = '中‘ ❌ 报错，正确如霞：
	var c4 int = '中'
	//打印时选择类型,选择格式化进行输出
	fmt.Printf("c4是：%c\n", c4)

	fmt.Println("..........下列是关于转义字符.............")
	//练习转义字符
	// \n 换行
	fmt.Println("aaaa\nbbbb")
	// \b 退格
	fmt.Println("aaaa\bbbbb")
	// \r 光标回到本行开头，后续输入就会替换原有字符
	`
	fmt.Println(s4)

	//4.字符串拼接效果：
	var s5 string = "123" + "456"
	s5 += "789"
	fmt.Println(s5)

	//当一个字符串过长，注意 + 号保留在上一行的最后
	var s6 = "123" + "456" + "ABC" + "123" +
		"456" + "ABC" + "123" + "456" + "ABC" +
		"123" + "456" + "ABC" + "123" + "456" +
		"ABC" + "123" + "456" + "ABC"

	fmt.Println(s6)

}
