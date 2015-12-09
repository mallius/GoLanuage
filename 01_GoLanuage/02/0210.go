package main

import (
	"fmt"
)

func main() {
	var ch1 rune
	ch1 = '中'
	ch2 := 22269
	fmt.Println(ch1)
	fmt.Println(string(ch1) + string(ch2))
}
