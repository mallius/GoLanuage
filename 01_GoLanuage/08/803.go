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
	stu1 := student{13001, "liming", false, 19, "net01"}
	stu2 := student{13002, "zhangxiao", true, 18, "net01"}
	stu3 := student{name: "wangle", age: 19}
	stu4 := &student{name: "zhaoqiong", id: 13003}
	fmt.Println(stu1)
	fmt.Println(stu2)
	fmt.Println(stu3)
	fmt.Println(stu4)
}
