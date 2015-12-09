package main

import (
	"fmt"
)

func main() {
	var id int = 100
	var name string = "Liming"
	var grade float32 = 91.5

	fmt.Printf("%v %v %v\n", id, name, grade)
	fmt.Printf("%#v %#v %#v\n", id, name, grade)
	fmt.Printf("%T %T %T\n", id, name, grade)
	fmt.Printf("60% %\n")
}
