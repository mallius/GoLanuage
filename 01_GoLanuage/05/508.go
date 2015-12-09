package main

import (
	"fmt"
)

func main() {
	var i, sum = 1, 0
HERE:
	sum = sum + i
	i++
	if i <= 100 {
		goto HERE
	}
	fmt.Println(sum)
}
