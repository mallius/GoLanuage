package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const GOAL int = 36

func main() {
	var data int
	var count int
	var c byte
	var str string

	fmt.Println("输入一个正整数")
LABEL1:
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	/*
		键盘输入处理流程
	*/
	for i := 0; ; i++ {
		c, _ = r.ReadByte()
		if c == 10 || c == 13 {
			break
		} else {
			w.Flush()
			str += string(c)
		}
	}

	/*
		猜数字游戏
	*/

	for {
		data, _ = strconv.Atoi(str)
		if data > GOAL {
			fmt.Printf("数字%d 有点大\n", data)
			count++
			str = ""
			goto LABEL1
		} else if data < GOAL {
			fmt.Printf("数字%d 有点小\n", data)
			count++
			str = ""
			goto LABEL1
		} else {
			fmt.Println("正确")
			break
		}
	}
	fmt.Printf("count:%d\n", count)
}
