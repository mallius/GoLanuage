package main

import (
	"fmt"
)

func main() {
	f1()
	f2()
	fmt.Printf("\n")
	fmt.Println(f3())
}

func f1() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

//延迟顺序
//Printf(0)
//Printf(1)
//Printf(2)
//Printf(3)
//Printf(4)
//执行顺序反过来4,3,2,1,0
func f2() {
	for i := 0; i < 4; i++ {
		defer fmt.Printf("%d ", i)
	}
}

//函数有返回值，延迟执行的函数读取返回值后加1
func f3() (i int) {
	defer func() {
		i++
	}()
	return 1
}
