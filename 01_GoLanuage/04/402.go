package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scanln(&a, &b)
	if a > b {
		fmt.Println("a>b")
	}
	if (a != 0) && (a > 0) {
		fmt.Println("a is positive number")
	}
	if (a != 0) && (a < 0) {
		fmt.Println("a is negative number")
	}
	fmt.Println("OVER")
}
