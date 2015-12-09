package main

import (
	"fmt"
)

func main() {
	var map1 = map[string]int{"key1": 100, "key2": 200, "key3": 300}
	for k, v := range map1 {
		fmt.Println(k, v)
		if k == "key2" {
			delete(map1, k)
		}
		if k == "key3" {
			map1["key4"] = 400
		}
	}
	fmt.Println(map1)
}
