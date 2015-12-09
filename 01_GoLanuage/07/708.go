package main

import (
	"fmt"
)

func main() {
	f := closures(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

func closures(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
