package main

import (
	"fmt"
)

func main() {

	var num uint64

	fmt.Println("Input a num >= 0")
	fmt.Scanf("%d", &num)
	fmt.Println(fac(num))
	fmt.Println(fac_r(num))
}

func fac(n uint64) uint64 {

	var ret uint64 = 1
	var i uint64 = 1

	if n < 0 {
		fmt.Println("error n < 0")
	}
	if n == 0 {
		return 1
	}
	for i = 1; i <= n; i++ {
		ret = ret * i
	}
	return ret
}

//error
func fac_r(n uint64) (ret uint64) {
	if n < 0 {
		fmt.Println("error n < 0")
	}
	if n == 0 {
		ret = 1
	} else {
		ret = (n - 1) * 1
	}
	return
}
