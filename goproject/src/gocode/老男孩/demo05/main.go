package main

import "fmt"

func prinnum(start, end int) {
	for i := start; i <= end; i++ {
		fmt.Println(i)
	}
}

func add(x, y int) {
	fmt.Println(x + y)
}

func main() {
	fmt.Println("hello,world")
	prinnum(1, 5)
	fmt.Println("hello,world")

	add(10, 90)
	add(100, 900)

}
