package main

import (
	"fmt"
)

type People struct {
	Name string
	Age  int
}

type Tester interface{}

func main() {
	people := People{"zhangsan", 20}
	it := make([]Tester, 4)
	it[0], it[1], it[2], it[3] = 1, "Hello", people, true

	for index, element := range it {
		if value, ok := element.(int); ok {
			fmt.Printf("it[%d] is an int. value = %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("it[%d] is a string. value = %s\n", index, value)
		} else if value, ok := element.(People); ok {
			fmt.Printf("it[%d] is a People. value = %v\n", index, value)
		} else {
			fmt.Printf("it[%d] is an unknown type.\n", index)
		}
	}
}
