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

func main() {
}
