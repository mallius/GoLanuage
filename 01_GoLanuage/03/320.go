package main

import (
	"fmt"
)

func main() {
	var sum1, sum2, sum3 int
	var average float32
	fmt.Println("请输入3个整数")
	fmt.Scanf("%d,%d,%d", &sum1, &sum2, &sum3)
	average = float32(sum1+sum2+sum3) / 3.0
	fmt.Printf("aver=%6.2f\n", average)
}
