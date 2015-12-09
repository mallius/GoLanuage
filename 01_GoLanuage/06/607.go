package main

import (
	"fmt"
)

func main() {
	var slice1 = []int{1, 2, 3, 4, 5}
	var slice2 = make([]int, 5)
	var slice3 = make([]int, 5, 10)

	fmt.Printf("len = %2d cap = %2d %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("len = %2d cap = %2d %v\n", len(slice2), cap(slice1), slice2)
	fmt.Printf("len = %2d cap = %2d %v\n", len(slice3), cap(slice1), slice3)
}
