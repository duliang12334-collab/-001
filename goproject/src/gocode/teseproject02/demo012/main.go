package main

import (
	"fmt"
)

func main() {
	var scores [5]int
	for i := 0; i < len(scores); i++ {
		fmt.Printf("请录入第%d个学生的成绩", i+1)
		fmt.Scanln(&scores[i])
	}

	for i := 0; i < len(scores); i++ {
		fmt.Printf("第%d个学生的成绩: %d \n", i+1, scores[i])

	}
	for _, value := range scores {
		fmt.Printf("%d \n", value)
	}
}
