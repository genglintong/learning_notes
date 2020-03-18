package main

import "fmt"

func sum(vars ...int) int {
	sum := 0
	for _, i := range vars{
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1,2,3,4,5))
	fmt.Println(sum())
}
