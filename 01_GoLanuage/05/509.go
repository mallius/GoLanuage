package main

import (
	"fmt"
)

func main() {
	var str string = "Golang"
	for k, v := range str {
		fmt.Printf("s[%d]=%c\n", k, v)
	}
}
