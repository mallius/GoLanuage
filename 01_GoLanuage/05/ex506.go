package main

import (
	"fmt"
)

func main() {
	var height float64 = 100.00
	var half float64
	half = height
	for i := 0; i < 10; i++ {
		half = half / 2.00
		fmt.Println(i+1, half)
	}
}
