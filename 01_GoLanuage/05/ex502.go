package main

import (
	"fmt"
)

func main() {
	var arr = [10]int{1, 2, 3, 4, 22, 5, 100, 10000, 18000, 0}
	var ret int

	for i := 0; i < 10; i++ {
		if arr[i] > ret {
			ret = arr[i]
		}
	}
	fmt.Println(ret)
}
