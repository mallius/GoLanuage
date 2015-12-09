package main

import (
	"fmt"
)

func main() {

	var a, b, c int
	fmt.Println("请输入正整数a和b:")
	fmt.Scanf("%d %d", &a, &b)
	c = multiple(a, b)
	fmt.Printf("[%d][%d]：最小公倍数[%d]\n", a, b, c)
}

func divisor(a, b int) int {
	var r int
	for {
		r = a % b
		if r != 0 {
			a = b
			b = r
		} else {
			break
		}
	}
	return b
}

func multiple(a, b int) int {
	var d int
	d = divisor(a, b)
	return (a * b / d)
}
