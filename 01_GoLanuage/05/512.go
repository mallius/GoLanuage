package main

import (
	"fmt"
)

func main() {
	var str string

	var charCount, numCount, otherCount int
	fmt.Println("input strings...")
	fmt.Scanf("%s", &str)
	for _, v := range str {
		if v >= 48 && v <= 57 {
			numCount++
		} else if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
			charCount++
		} else {
			otherCount++
		}
	}

	fmt.Printf("%d\n", numCount)
	fmt.Printf("%d\n", charCount)
	fmt.Printf("%d\n", otherCount)
}
