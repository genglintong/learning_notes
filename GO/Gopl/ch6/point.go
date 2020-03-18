package main

import (
	"fmt"
	"image/color"
	"sync"
)

type Point struct {
	X int
	Y int
}

type ColorPoint struct {
	Point
	Color color.RGBA
}

func (p *Point) Distance() int {
	return p.X * p.X + p.Y * p.Y
}

var cache = struct{
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func lookUp(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()

	return v
}

func main() {
	var cp ColorPoint
	cp.X = 1
	cp.Y = 2

	fmt.Println(cp.Point.X)
	fmt.Println(cp.Distance())
}
