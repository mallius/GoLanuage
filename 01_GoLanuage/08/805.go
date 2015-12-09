package main

import (
	"fmt"
)

func main() {
	var usr1 = struct {
		id       int
		name     string
		password string
	}{name: "zhangsan"}
	fmt.Println(usr1)

	var usr2 = new(struct {
		id       int
		name     string
		password string
	})
	usr2.name = "lisi"
	fmt.Println(usr2)
}
