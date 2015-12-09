package main

import (
	"fmt"
)

func main() {
	var words string
	var ch byte
	var ln, count int

	fmt.Scanf("%s", &words)
	fmt.Println("\n")

	fmt.Scanf("%c", &ch)
	ln = len(words)

	for i := 0; i < ln; i++ {
		if byte(words[i]) == ch {
			count++
		}
	}
	fmt.Printf("There are %d %q in the %q.\n", count, ch, words)
}
