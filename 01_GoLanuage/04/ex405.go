package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scanln(&num)

	switch num {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("五")
	case 6:
		fmt.Println("六")
	case 7:
		fmt.Println("日")
	default:
		fmt.Println("error")
	}
}
