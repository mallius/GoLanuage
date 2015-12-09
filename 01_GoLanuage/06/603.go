package main

import (
	"fmt"
)

func main() {
	var f = [20]int{1, 1}
	for i := 2; i < 20; i++ {
		f[i] = f[i-2] + f[i-1]
	}
	for i := 0; i < 20; i++ {
		if i%5 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("f[%2d]=%4d ", i, f[i])
	}

}
