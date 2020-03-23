package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)
	n := 0

	go func() {
		for {
			chan1 <- "ping"
			str := <-chan2

			if str == "pong" {
				n++
			}
		}
	}()

	go func() {
		for str := range chan1 {
			if str == "ping" {
				chan2 <- "pong"
			}
		}
	}()

	for t := 10; t >= 0; t-- {
		time.Sleep(time.Second)
		fmt.Println(n)
	}
}
