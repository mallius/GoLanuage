package main

import (
	"fmt"
)

type d1 struct {
	x int
}

type d2 struct {
	x int
}

type data struct {
	d1
	d2
}

func main() {
	d := data{d1{10}, d2{20}}
	fmt.Println(d.d1.x, d.d2.x)
}
