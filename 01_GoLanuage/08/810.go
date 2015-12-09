package main

import (
	"fmt"
)

type people struct {
	name  string
	sex   bool
	phone string
}

type teacher struct {
	people
	department string
	phone      string
}

func main() {
	teacher1 := teacher{people{"郑智", false, "100-201"}, "Computer Science", "200-401"}
	fmt.Println(teacher1.phone)
	fmt.Println(teacher1.people.phone)
}
