package main

import (
	"fmt"
)

func main() {
	var arr = []int{12000, 2, 18000, 3, 2}
	var retMax, retMin int

	retMax, retMin = findMaxMin(arr)
	fmt.Println(retMax, retMin)
}

func findMaxMin(arr []int) (max, min int) {
	var imax, imin int

	imax = arr[0]
	imin = arr[0]
	for i, _ := range arr {
		if arr[i] > imax {
			imax = arr[i]
		}
		if arr[i] < imin {
			imin = arr[i]
		}
	}
	return imax, imin
}
