package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello, World!"
	var s1, s2, s3, s4 string = "llo", "go", "ll", "l"
	var ch1, ch2 rune = 'c', 'd'
	fmt.Printf("[%q] in [%q]? %t\n", s1, str, strings.Contains(str, s1))
	fmt.Printf("[%q] in [%q]? %t\n", s1, str, strings.Contains(str, s2))

	fmt.Printf("[%q] in [%q]? %t\n", s1, str, strings.ContainsAny(str, s1))
	fmt.Printf("[%q] in [%q]? %t\n", s1, str, strings.ContainsAny(str, s2))

	fmt.Printf("[%q] in [%q]? %t\n", s1, str, strings.ContainsRune(str, ch1))
	fmt.Printf("[%q] in [%q]? %t\n", s1, str, strings.ContainsRune(str, ch2))

	fmt.Printf("[%q]中含有[%d]个子串[%q]\n", str, strings.Count(str, s3), s3)
	fmt.Printf("[%q]中含有[%d]个子串[%q]\n", str, strings.Count(str, s4), s4)

}
