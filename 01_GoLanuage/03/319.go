package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = -1024
	var s string = "1024"
	value, _ := strconv.Atoi(s)
	str := strconv.Itoa(i)
	fmt.Printf("string[%q]-->int[%d]\n", s, value)
	fmt.Printf("int[%d]-->string[%q]\n", i, str)
}
