package main

import "fmt"

var x int = 0

func add(a int) {
	x += a
	fmt.Println("add" , a, x)
}

func sub(a int)  {
	x -= a
	fmt.Println("sub" , a , x)
}

func mul(a int)  {
	x *= a
	fmt.Println("mul" , a, x)
}

func main() {
	// 36
	defer sub(1)
	// 37
	defer sub(2)
	// 39
	defer sub(3)
	// 42
	defer mul(1)
	// 42
	defer mul(2)
	// 21
	defer mul(3)
	// 7
	defer add(1)
	// 6
	defer add(2)
	// 4
	defer add(3)

	x++
}
