package main

import (
	"fmt"
)

type People struct {
	name string
	id   int
}

func main() {
	var people People = People{"张三", 12}
	var Slice []interface{}
	Slice = make([]interface{}, 8)

	Slice = []interface{}{1, "hello", people, people, people, people, 2, 3}

	for index, element := range Slice {
		switch value := element.(type) {
		case int:
			fmt.Printf("Slice[%d] is a int. value = %d\n", index, value)
		case string:
			fmt.Printf("Slice[%d] is a string. value = %s\n", index, value)
		case People:
			fmt.Printf("Slice[%d] is a People. value = %v\n", index, value)
		default:
			fmt.Printf("Slice[%d] is unknown. \n", index)
		}
	}
}
