package main

import "fmt"

// 数据类型默认值
func main() {
	var a int
	var b float32
	var c float64
	var d bool
	var e string

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Println("-----------下列是类型转换----------")
	//进行类型转换
	var n1 int = 100
	//var n2 float32 = n1 在这里自动转换不好使，需要显示转换,例如：
	var n2 float32 = float32(n1)
	fmt.Println(n1)
	fmt.Println(n2)
	//注意：n1的类型还是int类型，只是将值100转换为float32而已，如下：
	fmt.Printf("%T", n1) // 输出类型还是int
	fmt.Println()

	//将int64转为 int8的时候，编译不会出错，但是数据会溢出
	var n3 int64 = 888888
	var n4 int8 = int8(n3)
	fmt.Println(n4) //56 in8无法显示出完整的int64位，所以溢出了

	var n5 int32 = 12
	var n6 int64 = int64(n5) + 30 //一定要匹配 = 左右的数据类型
	fmt.Println(n6)

	//特殊情况
	var n7 int64 = 12
	var n8 int8 = int8(n7) + 127 // 编译用过，但结果可能会溢出
	// var n9 int8 = int(n7) + 128 //编译不用过
	fmt.Println(n8)

}
