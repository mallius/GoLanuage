package main

import (
	"fmt"
)

type skills []string

type people struct {
	name string
	sex  bool
}

type teacher struct {
	people
	skills
	int
	department string
}

func main() {
	teacher1 := teacher{}
	teacher1.name = "郑智"
	teacher1.skills = append(teacher1.skills, "Computer", "Golang")
	teacher1.int = 100
	teacher1.department = "Computer Science"
	fmt.Println(teacher1)
}
