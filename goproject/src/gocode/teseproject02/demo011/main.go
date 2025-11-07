package main

import (
	"errors"
	"fmt"
)

func main() {
	err := test()
	if err != nil {
		fmt.Println("错误： ", err)
		panic(err)
	}
	fmt.Println("流程成功")
	fmt.Println("正常执行")

}

func test() (err error) {
	num1 := 10
	num2 := 0
	if num2 == 0 { //如果等于0
		return errors.New("该值不能为0") // 通过errors.New 返回自定义函数
	} else { // 反之正常执行
		result := num1 / num2
		fmt.Println(result)
		return nil //返回0值

	}

}
