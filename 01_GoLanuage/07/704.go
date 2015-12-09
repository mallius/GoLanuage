package main

import (
	"fmt"
)

func main() {
	f1(1, 2, 3)
	f1(1, 2, 3, 4, 5)
}

func f1(args ...int) {
	fmt.Println(args)
}
