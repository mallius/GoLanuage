package main

import (
	"fmt"
)

func main() {
	var a int
	var f float32
	var str string

	fmt.Println("ready to scan")
	fmt.Scan(&a, &f, &str)
	fmt.Println("result:")
	fmt.Println(a, f, str)
}
