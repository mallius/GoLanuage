package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var num int
	buf := make([]byte, 1024)
	f, _ := os.Open("golang.txt")
	defer f.Close()
	n, _ := f.Read(buf)
	fmt.Println(string(buf[:n]), n)
	for _, sentence := range bytes.FieldsFunc(buf[:n], f1) {
		for _, words := range bytes.Fields(sentence) {
			num++
			fmt.Printf("%q", words)
		}
	}
	fmt.Println()
	fmt.Println("文件", f.Name(), "共有", num, "单词")
}

func f1(a rune) bool {
	if a == ',' || a == '.' || a == '!' {
		return true
	} else {
		return false
	}
}
