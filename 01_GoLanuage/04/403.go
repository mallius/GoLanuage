package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanln(&a, &b)

	if a > b {
		fmt.Printf("MAX=%d\n", a)
	} else {
		fmt.Printf("MAX=%d\n", b)
	}
}
