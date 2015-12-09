package main

import (
	"fmt"
)

func main() {
	var a, b int
	var op byte
	fmt.Scanf("%d %c %d", &a, &op, &b)
	switch op {
	case '+':
		fmt.Println(a + b)
	case '-':
		fmt.Println(a - b)
	case '*':
		fmt.Println(a * b)
	default:
		fmt.Println(a / b)
	}

}
