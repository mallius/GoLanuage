package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scanf("%d", &i)
	switch i {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
	}
}
