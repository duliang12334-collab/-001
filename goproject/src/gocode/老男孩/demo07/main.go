package main

import (
	"fmt"
	"time"
)

func foo() {
	fmt.Println("foo功能开始")
	time.Sleep(time.Second * 2)
	fmt.Println("foo功能结束")
}

func bar() {
	fmt.Println("bar功能开始")
	time.Sleep(time.Second * 2)
	fmt.Println("bar功能结束")
}

func timer(f func()) {
	start := time.Now().Unix()
	f()
	end := time.Now().Unix()
	fmt.Println("timer:", end-start)
}

func main() {
	timer(foo)
	time.Sleep(time.Second * 3)
	timer(bar)

}
