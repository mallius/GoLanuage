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
	var maxScore int = 0
	var index int = 0

	/* 切片声明、赋值 */
	var stuSlice []student
	stuSlice = []student{
		{1, "stu1", 100},
		{2, "stu2", 500},
		{3, "stu3", 300},
	}
	fmt.Println(stuSlice)

	/* 遍历切片调用方法获取数据 */
	for i, _ := range stuSlice {
		if stuSlice[i].getScore() > maxScore {
			maxScore = stuSlice[i].getScore()
			index = i
		}
	}
	fmt.Println(maxScore)
	fmt.Println(stuSlice[index])

}

func (recv student) getScore() (score int) {
	return recv.score
}
