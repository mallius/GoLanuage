package main

import (
	"fmt"
)

func main() {
	var f, c, k float64
	fmt.Println("输入华氏温度:")
	fmt.Scanf("%f", &f)

	c = 5.00 / 9.00 * (f - 32)
	k = 273.16 + c
	fmt.Printf("c=%6.2f k=%6.2f\n", c, k)
}
