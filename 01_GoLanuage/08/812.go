package main

import (
	"fmt"
)

type people struct {
	name string
	sex  bool
}

type teacher struct {
	*people    //匿名类型指针
	department string
}

func main() {
	teacher1 := teacher{&people{"郑智", false}, "Computer Science"}
	//匿名类型指针访问和普通匿名字段访问方式相同
	fmt.Println(teacher1)
	fmt.Println(teacher1.name, teacher1.sex, teacher1.department)
}
