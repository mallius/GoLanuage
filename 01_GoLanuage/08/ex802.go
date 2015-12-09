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
	//var arrStu [2]student = [2]student{{1, "stu1", 100}, {2, "stu2", 200}}

	//var arrStu [2]student
	//arrStu = [2]student{{1, "stu1", 100}, {2, "stu2", 200}}

	arrStu := [4]student{{1, "stu1", 100}, {2, "stu2", 200}, {4, "stu4", 400}, {3, "stu3", 300}}

	findMax(arrStu)

	fmt.Println(arrStu)
}

/* 最好使用切片 */
func findMax(stu [4]student) {
	var maxScore int = stu[0].score
	var index int = 0

	for i, _ := range stu {
		if stu[i].score > maxScore {
			maxScore = stu[i].score
			index = i
		}
	}
	fmt.Println("成绩最好", stu[index])
}
