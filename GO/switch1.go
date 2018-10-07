package main

import (
	"fmt"
)

/**
* switch 接受任意形式的表达式
* 不需要break退出 匹配到分支后执行完自动退出
* 如需要继续进行 通过fallthrogh 达到目的
*/

func main() {
	num1 := 7

	switch {
		case num1 < 0 :
			fmt.Printf("nagative\n")
		case num1 > 0 && num1 < 10:
			fmt.Printf("between 0 and 10 \n")
		default:
			fmt.Printf("greater\n")
	}
}
