package main

import (
	"fmt"
)

type student struct {
	Id     int
	Name   string
	Sex    bool
	Scores [3]int
}

func main() {
	var s1 student
	s1.Id = 1
	s1.Name = "stu1-Zhao"
	s1.Sex = true
	s1.Scores[0] = 100
	s1.Scores[1] = 100
	s1.Scores[2] = 100

	s2 := student{}
	s2.Id = 2
	s2.Name = "stu2-Qian"

	fmt.Println(s1)
	fmt.Println(s2)
}
