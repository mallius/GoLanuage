package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	var res float64
	fmt.Println("input 3 num")
	fmt.Scanf("%f %f %f", &a, &b, &c)
	res = math.Max(a, b)
	res = math.Max(res, c)
	fmt.Println(res)
}
