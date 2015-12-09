package main

import (
	"fmt"
)

func main() {
	f1(1, 2, 3, 4, 5)
}

func f1(args ...int) {
	f2(args...)
	f2(args[2:]...)
}

func f2(args ...int) {
	fmt.Println(args)
}
