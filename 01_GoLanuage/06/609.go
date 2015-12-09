package main

import (
	"fmt"
)

func main() {
	var s1 = make([]int, 3, 6)
	s2 := append(s1, 1, 2, 3)
	s3 := append(s1, 1, 2, 3, 4)
	fmt.Printf("len = %d cap = %d %v\n", len(s1), cap(s1), s1)
	fmt.Printf("len = %d cap = %d %v\n", len(s2), cap(s2), s2)
	fmt.Printf("len = %d cap = %d %v\n", len(s3), cap(s3), s3)
}
