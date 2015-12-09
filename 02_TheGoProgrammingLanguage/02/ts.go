package main

import "fmt"

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func main() {
	var v1 bool
	v1 = true
	fmt.Println(v1)

	var value2 int32
	value1 := 64
	value2 = (int32)(value1)
	fmt.Println(value2)

	var fvalue1 float32
	fvalue1 = 12
	fvalue2 := 12.0

	fmt.Println(fvalue2, fvalue1)

	var str string
	str = "Hello world"
	ch := str[0]
	fmt.Println(str, ch)
	fmt.Printf("%c\n", ch)

	str1 := "Hello,世界"
	n := len(str1)
	for i, ch := range str1 {
		fmt.Println(i, ch)
	}
	fmt.Println(str1, n)

	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}

	var ai = [3]int{1, 2, 3}
	var b = &ai
	b[1]++
	fmt.Println(ai, *b)

}
