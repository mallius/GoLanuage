package main

import (
	"fmt"
)

type rectangle struct {
	width  int
	height int
}

func (recv rectangle) area() int {
	return recv.width * recv.height
}

func main() {
	r1 := rectangle{4, 3}
	r2 := rectangle{30, 15}
	fmt.Println(r1.area(), r2.area())
}
