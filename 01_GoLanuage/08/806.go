package main

import (
	"fmt"
)

func main() {
	map1 := map[string]struct {
		name string
		age  int
	}{
		"teacher": {"zhenzhi", 39},
		"student": {"liming", 18},
	}
	fmt.Println(map1)
}
