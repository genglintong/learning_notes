package main

import (
	"fmt"
)

func Multiply(a, b int, replay *int) {
	*replay = a * b
}

func main() {
	n := 0
	replay := &n
	Multiply(10, 5, replay)
	fmt.Println("%d" , *replay)
}
