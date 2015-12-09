package main

import (
	"fmt"
)

func main() {
	var sum1, sum2 int

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum1 += i
		} else {
			sum2 += i
		}
	}
	fmt.Println(sum1, sum2)
}
