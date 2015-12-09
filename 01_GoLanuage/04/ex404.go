package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("input a number")
	fmt.Scanln(&number)

	if (number%3 == 0) || (number%5 == 0) {
		fmt.Println("OK")
	}
}
