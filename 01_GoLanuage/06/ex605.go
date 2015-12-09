package main

import (
	"fmt"
)

func main() {
	map1 := map[string]int{
		"psw1": 123,
		"psw2": 456,
		"psw3": 789,
	}

	fmt.Println(map1)
	for k, _ := range map1 {
		if k == "psw2" {
			delete(map1, k)
		}
	}
	fmt.Println(map1)
	map1["psw4"] = 135 //增加
	fmt.Println(map1)
}
