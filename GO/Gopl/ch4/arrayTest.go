package main

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

	fmt.Println(symbol[USD])

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Println(PtrCount(&c1, &c2))

	fmt.Printf("%x\n%x\n%t\n%T", c1, c2, c1 == c2, c1)
}

func PtrCount(ptr1 *[32]uint8, ptr2 *[32]uint8) int {
	var count = 0
	for i := range ptr1{
		t := ptr1[i]^ptr2[i]
		for i := 0; i < 8; i++ {
			if t & (1 << i) > 0 {
				count++
			}
		}
	}
	return count
}
