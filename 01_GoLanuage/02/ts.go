package main

import (
	"fmt"
)

const (
	a = 'A'
	b = a + iota
	c
	d = iota
)

func main() {
	var a, b int = 10, 1
	var p *int
	p = &a
	a = *p + b
	fmt.Println(a)

	cp := 1.23 + 4.56i
	fmt.Println(cp)

	var str1 string
	var str2 string
	fmt.Scanf("%s%s", &str1, &str2)
	fmt.Println(str1 + str2)

	fmt.Println(a, b, c, d)
}
