package main

import (
	"fmt"
)

func main() {
	var map1 map[string]int
	map1 = map[string]int{"key1": 1, "key2": 2, "key3": 3}

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	var map2 map[string]interface{}

	map2 = map[string]interface{}{"key10": 10, "key11": 11, "key12": 12}
	for k, v := range map2 {
		fmt.Println(k, v)
	}
}
