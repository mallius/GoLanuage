package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "na"
	var count int = 2

	fmt.Println("ba" + strings.Repeat(s, count))
	fmt.Println(strings.Replace("google", "o", "oo", 1))
	fmt.Println(strings.Replace("google", "o", "oo", -1))
}
