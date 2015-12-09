package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scanf("%s", &str)
	for _, v := range str {
		fmt.Printf("%q ", v)
	}
}
