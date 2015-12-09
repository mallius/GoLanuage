package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, l, hl, area float64
	fmt.Println("请输入三角形三边")
	fmt.Scanf("%f,%f,%f", &a, &b, &c)
	l = a + b + c
	fmt.Printf("l=%6.f\n", l)
	hl = l * 0.5
	area = math.Sqrt(hl * (hl - a) * (hl - b) * (hl - c))
	fmt.Printf("area=%6.2f\n", area)
}
