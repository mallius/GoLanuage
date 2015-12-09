package main

import (
	"fmt"
)

type people struct {
	name  string
	phone string
}

type teacher struct {
	people
	department string
}

type student struct {
	people
	school string
}

func (r people) sayHi() {
	fmt.Printf("Hi, I'm %s you can call me on %s.\n", r.name, r.phone)
}
func main() {
	teacher1 := teacher{people{"郑智", "010-22002"}, "CS"}
	student1 := student{people{"黎明", "010-11002"}, "Yale"}
	teacher1.sayHi()
	student1.sayHi()
}
