package main

import (
	"fmt"
)

type people struct {
	name string
	sex  bool
}

type teacher struct {
	people
	department string
}

func main() {
	teacher1 := teacher{people{"郑智", false}, "Computer Science"}
	var teacher2 teacher = teacher{people{"郑智2", false}, "Computer Science"}

	fmt.Println(teacher1)
	fmt.Println(teacher2)
}
