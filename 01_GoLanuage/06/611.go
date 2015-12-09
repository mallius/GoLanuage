package main

import (
	"fmt"
)

func main() {
	var map1 map[string]int
	fmt.Println(map1)
	map1 = make(map[string]int)
	map1["key1"] = 1
	fmt.Println(map1)

	var map2 = map[string]int{}
	map2["key2"] = 2
	fmt.Println(map2)
}
