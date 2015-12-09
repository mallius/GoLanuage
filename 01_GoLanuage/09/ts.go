package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Name string
}

type Student struct {
	People
	School string
}

type Teacher struct {
	People
	Department string
}

func (p People) SayHi()  {}
func (s Student) SayHi() {}
func (s Student) Study() {}
func (t Teacher) SayHi() {}

type Speaker interface {
	SayHi()
}

type Learner interface {
	SayHi()
	Study()
}

func main() {
	var pi float64 = 3.14
	t := reflect.TypeOf(pi)
	v := reflect.ValueOf(pi)
	fmt.Println("Type:", t.Kind())
	fmt.Println("Value:", v.Float())
}
