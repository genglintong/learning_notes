package main

import (
	"fmt"
	"./pack1/pack1"
)

func main() {
	var test1 string
	test1 = pack1.ReturnStr()

	fmt.Printf("ReturnStr from package1 : %s \n", test1)
	fmt.Printf("Interger from package1: %d \n",pack1.Pack1Int)
}
