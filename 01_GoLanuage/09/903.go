package main

import (
	"fmt"
)

type People struct {
	Name string
	Age  int
}

type Student struct {
	People
	School string
}

type PeopleInfo interface {
	GetPeopleInfo()
}

type StudentInfo interface {
	GetPeopleInfo()
	GetStudentInfo()
}

func (p People) GetPeopleInfo() {
	fmt.Println(p)
}

func (s Student) GetStudentInfo() {
	fmt.Println(s)
}

func main() {
	var is StudentInfo = Student{People{"Liming", 18}, "Yale University"}
	is.GetStudentInfo()
	is.GetPeopleInfo()

	var ip PeopleInfo = is
	ip.GetPeopleInfo()
}
