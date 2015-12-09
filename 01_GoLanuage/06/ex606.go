package main

import (
	"fmt"
)

func main() {
	var slice = []float64{1.1, 2.0, 3.0, 4.0, 5.0}
	var tot, aver float64
	var num float64 = 0.0

	for _, v := range slice {
		tot += v
		num++
	}
	aver = tot / num
	fmt.Println(slice)
	fmt.Printf("tot:%v aver:%v\n", tot, aver)
}
