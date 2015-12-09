package main

import (
	"fmt"
)

func main() {
	var searchstr string
	fmt.Scanf("%s", &searchstr)
	switch searchstr {
	case "c", "C":
		fmt.Println("C programing language")
	case "go", "golang":
		fmt.Println("Go programing language")
	case "java":
		fmt.Println("Java")
	default:
		fmt.Println("Not find the result!")
	}
}
