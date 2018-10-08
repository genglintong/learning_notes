package main

import (
	"fmt"
)

func main() {
	/**
	* new 分配内存
	* make 切片 map channel 初始化
	*/

	// 初始化
	var slice1 []int = make ([]int, 10)

	for i := 0; i < len(slice1) ; i++ {
		slice1[i] = 5 * i
	}

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}
