package main

import (
	"fmt"
)

func main() {
	var str string
	str = "Go" + "lang"
	fmt.Println(str)
	fmt.Println("字符串长度:", len(str))
	fmt.Println("%v", str[1])
}
