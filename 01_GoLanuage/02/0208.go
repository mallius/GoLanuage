package main

import (
	"fmt"
)

func main() {
	var cp1 complex64
	cp1 = 1.2 + 3.4i
	fmt.Println("real:", real(cp1))
	fmt.Println("imag:", imag(cp1))
}
