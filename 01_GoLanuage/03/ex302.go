package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, sum, aver, sqt, sqt_sum float64
	fmt.Println("输入三个数:")
	fmt.Scanf("%f,%f,%f", &a, &b, &c)
	sum = a + b + c
	aver = sum / 3.0
	sqt_sum = a*a + b*b + c*c
	sqt = math.Sqrt(sqt_sum)
	fmt.Printf("sum=%6.2f aver=%6.2f\n", sum, aver)
	fmt.Printf("sqt_sum=%6.2f\n", sqt_sum)
	fmt.Printf("sqt=%6.2f\n", sqt)
}
