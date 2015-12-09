package main

import (
	"fmt"
)

func main() {
	var slice = []float64{1.0, 2.0, 3.3}
	fmt.Println(SliceAver(slice))
}

func SliceAver(slice []float64) float64 {
	var aver, sum float64 = 0.0, 0.0
	var num float64 = 0.0

	for i, _ := range slice {
		sum += slice[i]
		num++
	}
	aver = sum / num
	return aver
}
