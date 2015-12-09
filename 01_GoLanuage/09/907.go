package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Id    int
	Name  string
	Sex   bool
	Grade float32
}

func (s Student) SayHi() {
	fmt.Println("Hi, nice to meet you!\n")
}

func (s Student) MyName() {
	fmt.Println("My name is %s, I come from China.", s.Name)
}

//反射处理函数
func StructInfo(o interface{}) {
	t := reflect.TypeOf(o)
	if k := t.Kind(); k != reflect.Struct {
		fmt.Printf("This value is not a struct, it's %v.", k)
		return
	}
	fmt.Println("Struct name is:", t.Name())
	fmt.Println("Fields of the struct is:")
	v := reflect.ValueOf(o)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		fmt.Printf(" %6s: %v = %v\n", field.Name, field.Type, value)
	}
	fmt.Println("Methods of the struct is:")

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf(" %6s:%v\n", method.Name, method.Type)
	}
}

func main() {
	stu := Student{10001, "李敏", false, 90.5}
	StructInfo(stu)
}
