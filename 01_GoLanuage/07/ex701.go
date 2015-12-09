package main

import (
	"fmt"
	"math"
)

func main() {
	cal()
}
func cal() {
	var a, b, c float64
	var area float64
	var s float64

	fmt.Println("input a b c")
	fmt.Scanf("%f %f %f", &a, &b, &c)
	s = 0.5 * (a + b + c)

	area = math.Sqrt(s * (s - a) * (s - b) * (s - c))
	fmt.Println(area)

}
