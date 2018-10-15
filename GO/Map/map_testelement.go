package main

import (
	"fmt"
)

func main() {
	//var value int
	//var isPresent bool

	map1 := make(map[string]int)
	map1["New Delhi"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25
	if _, isPresent := map1["Beijing"]; isPresent {
		delete(map1,"Beijing")
		//fmt.Println(value)
		fmt.Println(isPresent)
	}
}
