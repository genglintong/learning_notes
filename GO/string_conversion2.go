package main

import (
	"fmt"
	"strconv"
)

func main() {
	orig := "123"
	var newS string

	fmt.Printf("int 长度：%d \n", strconv.IntSize)

	// 转int
	an, err := strconv.Atoi(orig)

	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error \n", orig)
		return
	}

	fmt.Printf("The integer is %d\n", an)
	an += 5
	newS = strconv.Itoa(an)
	fmt.Printf("the New string is :%s \n", newS)
}
