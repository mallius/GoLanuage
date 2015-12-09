package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2, str3 string = "GO", "go", "lang"
	fmt.Printf("%q == %q? %t\n", str1, str2, strings.EqualFold(str1, str2))
	fmt.Printf("%q == %q? %t\n", str1, str3, strings.EqualFold(str1, str3))
}
