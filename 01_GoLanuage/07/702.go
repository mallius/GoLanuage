package main

import (
	"errors"
	"fmt"
)

func f1() {
	fmt.Println("hello world!")
}

func f2(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func f3(a int, b int) int {
	return a + b
}

func f4(a int, b int) (ret float32, err error) {
	if b == 0 {
		err = errors.New("Overflow!")
		return
	} else {
		return float32(a) / float32(b), nil
	}
}

func main() {
	var a, b int = 2, 0
	var c string = "Golang"
	f1()
	f2(a, b, c)
	sum := f3(a, b)
	fmt.Println(sum)
	f, err := f4(a, b)
	fmt.Println(f, err)
}
