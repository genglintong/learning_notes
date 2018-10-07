package main

import (
	"fmt"
)

func main() {
	str := "GO is a beautiful language!"
	fmt.Printf("字符串长度为: %d\n", len(str))

	for i := 0 ; i < len(str) ; i++ {
		fmt.Printf("Character on position %d is : %c \n", i , str[i])
	}

	str = "日本语"
	fmt.Printf("字符串长度为: %d\n", len(str))

	for pos , char := range str {
		fmt.Printf("Character on position %d is : %c \n", pos , char)
	}
    for i := 0 ; i < len(str) ; i++ {
        fmt.Printf("Character on position %d is : %c \n", i , str[i])
    }
}
