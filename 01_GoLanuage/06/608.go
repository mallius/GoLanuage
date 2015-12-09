package main

import (
	"fmt"
)

func main() {
	var slice1 = []int{1, 2, 3, 4, 5}

	for i := 0; i <= 4; i++ {
		fmt.Printf("slice1[%d] = %d", i, slice1[i])
	}
	fmt.Printf("\n")

	for i, v := range slice1 {
		fmt.Printf("slice1[%d] = %d ", i, v)
	}
}
