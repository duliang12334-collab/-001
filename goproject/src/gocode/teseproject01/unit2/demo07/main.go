package main

import (
	"fmt"
	"strconv"
)

// 关于基本数据类型转为string
func main() {
	//方式1:fmt.Sprintf
	var n1 int = 18
	var n2 float32 = 1.23
	var n3 bool = false
	var n4 byte = 'A'

	var s1 string = fmt.Sprintf("%d", n1)
	fmt.Printf("s1类型: %T , s1 = %q \n", s1, s1)
	//s1类型: string , s1 = "18"

	var s2 string = fmt.Sprintf("%f", n2)
	fmt.Printf("s2类型: %T , s2 = %q \n", s2, s2)
	//s2类型: string , s2 = "1.230000"

	var s3 string = fmt.Sprintf("%t", n3)
	fmt.Printf("s3类型: %T , s3 = %q \n", s3, s3)
	//s3类型: string , s3 = "false"

	var s4 string = fmt.Sprintf("%c", n4)
	fmt.Printf("s4类型: %T , s4 = %q \n", s4, s4)
	//s4类型: string , s4 = "A"

	// 方式2: strconv.Format..(..表示对应类型)//比较麻烦，如下：
	var n5 int = 88
	var s5 string = strconv.FormatInt(int64(n5), 10) //参数：第一个表示必须转为int64位，第二个参数指定字面值为十进制
	fmt.Printf("s5类型: %T , s5 = %q \n", s5, s5)

	//string类型转 基本数据类型
	fmt.Println("---------------以下是string类型转 基本类型------------")
	//使用Parse..进行转换
	// string-->bool
	var x1 string = "trun"
	var b bool
	//PaeseBool这个函数的返回值有两个（value bool，err error)
	//value就是我们需要得到的数据，err是出现的错误
	//我们只关注得到的布尔类型数据，err 可以用 _ 直接忽略
	b, _ = strconv.ParseBool(x1)
	fmt.Printf("b类型: %T, b = %v \n", b, b)

	//string--->int64
	var x2 string = "28"
	var num1 int64
	num1, _ = strconv.ParseInt(x2, 10, 64)
	fmt.Printf("num1类型 %T, num1 = %v \n", num1, num1)

	//string--->float32/float64
	var x3 string = "3.1415"
	var f23 float64
	f23, _ = strconv.ParseFloat(x3, 64)
	fmt.Printf("f23类型 %T, f23 = %v \n", f23, f23)

	// string向基本类型转换时，一定要确定string类型时有效的，否则输出结果为对应结果的默认值

}
