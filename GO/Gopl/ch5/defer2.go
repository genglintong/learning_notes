package main

import (
	"fmt"
	"os"
	"runtime"
)

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	_, _ = os.Stdout.Write(buf[:n])
}

func recover2()  {
	if p := recover(); p != nil {
		_, _ = fmt.Printf("internal error: %v", p)
	}
}

func main() {
	defer recover2()
	f(3)
}
