package main

import (
	"fmt"
)

func main() {
	var array1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice1 []int
	slice1 = array1[:5]
	slice2 := array1[5:]
	slice3 := array1[4:7]
	slice4 := array1
	slice5 := array1[:]
	slice6 := array1[0:len(array1)]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(slice6)
}
