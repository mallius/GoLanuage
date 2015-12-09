package main

import (
	"fmt"
)

func main() {
	var psw [10]int
	var num, m, n, k int
	fmt.Println("请输入总人数n 和初始值m")
	fmt.Scan(&n, &m)

	fmt.Println("请输入密码初始化数组")
	for i := 0; i < n; n++ {
		fmt.Scan(&psw[i])
	}
	fmt.Println("密码初始化完毕，开始报数...")
Label1:
	for {

	}

}
