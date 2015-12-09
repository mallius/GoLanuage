package main

import (
	"fmt"
	"math"
)

func main() {
	var num, i, k int

	fmt.Scanf("%d", &num)
	k = int(math.Sqrt(float64(num)))
	for i = 2; i <= k; i++ {
		if num%i == 0 {
			break
		}
	}
	if i > k {
		fmt.Printf("The number %d is a primer number.\n", num)
	} else {
		fmt.Printf("the number %d is not a prime number.\n", num)
	}

}
