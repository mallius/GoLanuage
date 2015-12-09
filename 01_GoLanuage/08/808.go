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

type departmentHead struct {
	teacher
	college string
}

func main() {
	head := departmentHead{}
	head.name = "郑智"
	head.department = "Computer science"
	head.college = "Yale University"
	fmt.Println(head)
}
