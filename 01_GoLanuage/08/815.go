package main

import (
	"fmt"
)

type object struct {
	id   int
	name string
}

func (object) msgbox() {
	fmt.Println("This is a object!")
}

func (*object) msgBox() {
	fmt.Println("This is a * object!")
}

func main() {
	obj := object{}
	p := &obj
	obj.msgbox()
	p.msgBox()
}
