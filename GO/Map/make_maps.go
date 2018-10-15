package main

import "fmt"

func main() {
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one":1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.1415
	mapAssigned["two"] = 3

	fmt.Println(mapCreated)
	fmt.Println(mapAssigned)
	fmt.Println("vim-go")
}
