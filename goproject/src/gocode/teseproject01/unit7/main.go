package main

import "fmt"

func main() {
	// 定义一个字符串
	var str string = "hello golang 你好"
	// 方式1:普通 for 循环：按照字节进行遍历输出
	//for i := 0; i < len(str); i++ {
	//	fmt.Printf("%c \n", str[i])
	//}

	// 方法2:for range
	for i, value := range str {
		fmt.Printf("索引为： %d,具体的值为： %c \n", i, value)
	}

	// 对 str 进行遍历，遍历的每个解说索引值被i接收，每个结果的具体值被value接收

}
