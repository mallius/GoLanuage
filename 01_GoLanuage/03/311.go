package main

import (
	"fmt"
)

func main() {
	var a int
	var f float32
	var str string

	fmt.Scanln(&a, &f, &str)
	fmt.Println(a, f, str)
}
