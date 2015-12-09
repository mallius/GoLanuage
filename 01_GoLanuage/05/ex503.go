package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	var numstr string
	/* 水仙花数三重循环 */
	for i := 1; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				numstr = strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(k)
				num, _ = strconv.Atoi(numstr)
				if num == (i*i*i)+(j*j*j)+(k*k*k) {
					fmt.Println(num)
				}
			}
		}
	}

}
