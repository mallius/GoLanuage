package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, disc, x1, x2, p, q float64
	var cpx1, cpx2 complex128

	fmt.Println("输入一元二次方程的三个系数")
	fmt.Scanf("%f,%f,%f", &a, &b, &c)

	if a == 0 {
		fmt.Println("系数a==0不是一元二次方程")
	} else {
		disc = b*b - 4*a*c
		p = -b / (2 * a)
		q = math.Sqrt(disc) / (2 * a)
		if disc == 0 {
			x1 = p
			x2 = q
			fmt.Printf("x1=x2=%8.2f\n", p)
		} else if disc > 0 {
			x1 = p + q
			x2 = p - q
			fmt.Printf("x1=%8.2f x2=%8.2f\n", x1, x2)
		} else if disc < 0 {
			q = math.Sqrt(math.Abs(disc)) / (2 * a)
			cpx1 = complex(p, q)
			cpx2 = complex(p, -q)
			fmt.Printf("cpx1=%8.2f cpx2=%8.2f\n", cpx1, cpx2)
		}
	}

}
