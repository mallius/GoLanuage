package main

import (
	"fmt"
)

type student struct {
	id    int
	name  string
	score int
}

func main() {
	var stuSlice []student
	stuSlice = []student{{1, "stu1", 100}, {2, "stu2", 200}, {3, "stu3", 300}}

	findMax(stuSlice)
	fmt.Println(stuSlice)
}

func findMax(slice []student) {
	var maxScore int = slice[0].score
	var index int = 0

	for i, _ := range slice {
		if slice[i].score > maxScore {
			maxScore = slice[i].score
			index = i
		}
	}
	fmt.Println(slice[index])
}
