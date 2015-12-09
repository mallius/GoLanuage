package main

import (
	"fmt"
)

func main() {
	var map1 = map[string]int{"key1": 100, "key2": 200}
	v, OK := map1["key1"]
	if OK {
		fmt.Println(v, OK)
	} else {
		fmt.Println(v)
	}

	v, OK = map1["key3"]
	if OK {
		fmt.Println(v, OK)
	} else {
		fmt.Println(v)
	}
}
