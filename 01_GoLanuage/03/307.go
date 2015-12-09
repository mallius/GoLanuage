package main

import (
	"fmt"
)

func main() {
	var str string = "Go language"

	fmt.Println("字符串和切片")

	fmt.Printf("%s\n", str)
	fmt.Printf("%q\n", str)
	fmt.Printf("%x\n", str)
	fmt.Printf("%X\n", str)
}
