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

type Student struct {
	People
	School string
}

func (p People) SayHi() {
	fmt.Printf("Hi, I'm %s. Nice to meet you!\n", p.Name)
}

func (t Teacher) SayHi() {
	fmt.Printf("Hi,my name is %s.I'm working in %s\n", t.Name, t.Department)
}

func (s Student) SayHi() {
	fmt.Printf("Hi,my name is %s. I'm studing in %s.\n", s.Name, s.School)
}

func (s Student) Study() {
	fmt.Printf("I'm learning Golang in %s.\n", s.School)
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
	student := Student{People{"郑智"}, "Yale University"}
	var is Speaker
	is = people
	is.SayHi()
	is = teacher
	is.SayHi()
	is = student
	is.SayHi()

	var il Learner
	il = student
	il.Study()

	ix := make([]Speaker, 3)
	ix[0], ix[1], ix[2] = people, teacher, student
	for _, value := range ix {
		value.SayHi()
	}
}
