package main

import (
	"fmt"
	"math"
)

func main() {
	var r, sita float64
	var x, y float64

	fmt.Println("输入极坐标")
	fmt.Scanf("%f,%f", &r, &sita)

	x = r * math.Cos(sita)
	y = r * math.Sin(sita)

	fmt.Printf("x=%6.2f y=%6.2f\n", x, y)

}
