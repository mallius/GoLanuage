package main

import (
	"fmt"
)

func main() {
	var s string
	var lenth int
	fmt.Println("请输入字符串")
	fmt.Scanf("%s", &s)
	lenth = len(s)
	fmt.Println(lenth, s)
}
