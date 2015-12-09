package main

import (
	"fmt"
)

type date struct {
	year  int
	month int
	day   int
}

type student struct {
	Id       int
	name     string
	sex      bool
	class    string
	birthday date
	_        string
}

func main() {
	stu := new(student)
	stu.name = "Liming"
	stu.birthday.year = 1995
	fmt.Println(stu)
}
