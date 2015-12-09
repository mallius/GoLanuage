package main

import (
	"fmt"
)

func main() {
	var num int
	for i := 1; i < 10; i++ {
		num += i
	}
	fmt.Println(num)
}
