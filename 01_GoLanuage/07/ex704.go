package main

import (
	"fmt"
)

func main() {
	fmt.Println(max(2.1, 3.1, 4.1))
}
func max(arg ...float64) float64 {
	var max float64
	max = arg[0]
	for _, v := range arg {
		if v > max {
		}
		max = v
	}
	return max
}
