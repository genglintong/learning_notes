package main

import "fmt"

// 匿名函数不仅仅是一段代码 还记录了状态
func squares() func() int {
	var x int

	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
