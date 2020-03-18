package main

import (
	"flag"
	"fmt"
	"time"
)

// flag.Duration 创建一个 time.Duration 类型的标记变量
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
