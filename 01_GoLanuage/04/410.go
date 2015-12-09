package main

import (
	"fmt"
)

func main() {
	var grade int
	fmt.Scanln(&grade)
	switch {
	case (grade >= 90) && (grade <= 100):
		fmt.Println("A")
	case (grade >= 80) && (grade < 90):
		fmt.Println("B")
	default:
		fmt.Println("D")

	}
}
