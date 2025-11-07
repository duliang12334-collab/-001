package main

import (
	"fmt"
)

func main() {

	var i, a int
	for i = 1; i <= 100; i++ {
		a += i
	}
	fmt.Println(a)

	// 退出循环
	for j := 0; j < 10; j++ {
		if j == 6 {
			continue
		}
		fmt.Println(j)
	}

}
