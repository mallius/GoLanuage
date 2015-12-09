package main

import (
	"fmt"
)

type student struct {
	Id       int
	name     string
	sex      bool
	class    string
	birthday struct {
		year  int
		month int
		day   int
	}
}

func main() {
	stu := new(student)
	stu.name = "liming"
	stu.birthday.year = 1995
	fmt.Println(stu)
}
