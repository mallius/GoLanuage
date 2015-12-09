package main

import (
	"fmt"
)

type Data struct {
	Num int
}

func (p *Data) Get() int {
	return p.Num
}

func (p *Data) Put(v int) {
	p.Num = v
}

type InOuter interface {
	Get() int
	Put(v int)
}

func main() {
	var data Data
	data = Data{100}

	var inout InOuter

	inout = &data
	fmt.Println(inout.Get())
	inout.Put(10)
	fmt.Println(inout.Get())
}
