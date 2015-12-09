package main

import (
	"fmt"
)

func main() {
	var i int = 100
	var i_pointer *int
	i_pointer = &i
	fmt.Printf("%p\n", i_pointer)
}
