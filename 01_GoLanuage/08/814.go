package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width  int
	height int
}

type circle struct {
	radius float32
}

func (recv rectangle) area() int {
	return recv.width * recv.height
}
func (recv circle) area() int {
	return recv.radius * recv.radius * math.Pi
}
func main() {
	r1 := rectangle{4, 3}
	r2 := rectangle{30, 15}
	fmt.Println(r1.area(), r2.area())
}
