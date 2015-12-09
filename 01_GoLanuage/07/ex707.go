package main

import (
	"fmt"
)

func main() {
	var str []byte

	defer func() {
		fmt.Println("defer func 处理非法字符")
		if err := recover(); err != nil {
			fmt.Println("defer func 字符串str含有非法字符", err)
		} else {
			fmt.Println("defer func 字符串中无非法字符")
		}
	}()

	fmt.Println("Input a string")
	fmt.Scanf("%s", &str)

	for _, v := range str {
		if v == ' ' || v == '!' || v == '#' {
			panic("main 字符串含有非法字符")
		}
	}
}
