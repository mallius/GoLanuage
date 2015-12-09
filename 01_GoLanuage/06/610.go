package main

import (
	"fmt"
)

func main() {
	var s1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var s2 = make([]int, 3, 5)
	var n int

	n = copy(s2, s1)
	fmt.Println(n, s2, len(s2), cap(s2))
	s3 := s1[3:6]
	n = copy(s3, s1[1:5])
	fmt.Println(n, s1, s3)
}
