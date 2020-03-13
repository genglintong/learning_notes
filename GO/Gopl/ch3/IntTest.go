package main

import "fmt"

// 2^1 + 2^5
var x uint8 = 1<<1 | 1<<5
// 2^1 + 2^2
var y uint8 = 1<<1 | 1<<2

func main() {
	// 以二进制方式输出
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	// 无符号数 只在位运算以及其他特殊场景使用
	for i := uint(0); i < 8; i++  {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}
