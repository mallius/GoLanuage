package main

import (
	"fmt"
)

func main() {
	var a int = 97
	var f float32 = -1.32
	var p *int
	var str string = "Golang"
	p = &a

	fmt.Printf("%+d, %+g\n", a, f)
	fmt.Printf("%+q\n", a)
	fmt.Printf("%d % #o % #x % #X\n", a, a, a, a)
	fmt.Printf("%p % #p\n", p, p)
	fmt.Printf("%s, % x\n", str, str)
}
