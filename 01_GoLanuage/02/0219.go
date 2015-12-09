package main

import (
	"fmt"
)

func main() {
	var str string
	str = "Hello世界!"
	n := len(str)
	fmt.Println("字节数组方式遍历:")
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Printf("str[%d] = %v\n", i, ch)
	}

	fmt.Println("Unicode")
	for i, ch := range str {
		fmt.Printf("str[%d] = %v\n", i, ch)
	}
}
