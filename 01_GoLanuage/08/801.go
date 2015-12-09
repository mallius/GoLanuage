package main

import (
	"fmt"
)

type student struct {
	id    int
	name  string
	sex   bool
	age   int
	class string
}

func main() {
	var stu1 student
	stu2 := student{}
	stu1.name = "Liming"
	stu2.name = "zhangheng"
	stu1.age = 18
	stu2.age = 19
	fmt.Println(stu1)
	fmt.Println(stu2)
}
