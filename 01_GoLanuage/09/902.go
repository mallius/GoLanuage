package main

import (
	"fmt"
)

type People struct {
	Name string
}

type Teacher struct {
	People
	Department string
}

func (p People) SayHi() {
	fmt.Printf("Hi, I'm %s. Nice to meet you!\n", p.Name)
}

type Speaker interface {
	SayHi()
}

type Learner interface {
	SayHi()
	Study()
}

func main() {
	people := People{"张三"}
	teacher := Teacher{People{"郑智"}, "Computer Science"}
	var is Speaker
	is = &people
	is.SayHi()
	is = &teacher
	is.SayHi()
}
