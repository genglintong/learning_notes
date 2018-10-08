package main

import (
	"fmt"
)

func main() {
	callback(1, Add)
}

func Add(a,b int) {
	fmt.Printf("The sum of %d and %d is %d \n",a, b, a+b)
}

// 回调函数
func callback(y int , f func(int,int)) {
	f(y,2)
}
