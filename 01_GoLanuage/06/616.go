package main

import (
	"fmt"
)

func main() {
	monthdays := map[string]int{
		"Jan": 31,
		"Feb": 28,
		"Mar": 31,
		"Apr": 30,
		"May": 31,
		"Jun": 30,
		"Jul": 31,
		"Aug": 31,
		"Sep": 30,
		"Oct": 31,
		"Nov": 30,
		"Dec": 31,
	}
	var year string
	fmt.Println("input leap or no leap")
	fmt.Scanf("%s", &year)
	if year == "leap" {
		monthdays["Feb"] = 29
	}
	totleDays := 0

	for _, days := range monthdays {
		totleDays += days
	}
	fmt.Printf("%d days\n", totleDays)
}
