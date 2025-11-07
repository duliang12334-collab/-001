package main

import "fmt"

func main() {
	// 功能：求1-100的和，当和第一次超过300的时候，停止程序
	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
		fmt.Println(sum)
		if sum >= 300 {
			//如果和大于等于300，停止这个循环（其他后续程序不受影响）
			break
		}

	}
	fmt.Println("123321")

	//双重循环：
	fmt.Println("---------以下是双循环示例----------")
lable1: // 加标签可以选择停止的循环对象
	for a := 1; a <= 5; a++ {
		for b := 5; b >= 1; b-- {
			fmt.Printf("a: %v,b: %v \n", a, b)
			if a == 3 && b == 3 {
				break lable1 // 使用标签停止指定的循环对象
			}

		}
	}

}
