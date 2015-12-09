package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y float64

	fmt.Println("输入x")
	fmt.Scanln(&x)

	if x >= 0 {
		y = (math.Sin(x) + math.Cos(x)) / 2
	} else {
		y = (math.Sin(x) - math.Cos(x)) / 2
	}
	fmt.Println(y)
}
