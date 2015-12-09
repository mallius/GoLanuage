package main

import (
	"fmt"
)

type d1 struct {
	x int
}

type d2 struct {
	x int
}

type dd struct {
	d1
	d2
}

type coordinate struct {
	x int
	y int
}

func (recv coordinate) swap() {
	var temp int
	temp, recv.x, recv.y = temp, recv.y, recv.x
	fmt.Println(recv)

}

func (recv *coordinate) ptrswap() {
	var temp int
	temp, recv.x, recv.y = temp, recv.y, recv.x
	fmt.Println(recv)
}

func main() {

	/* Swap */
	var r1 coordinate
	r1 = coordinate{3, 4}
	r1.swap()
	fmt.Println(r1)

	/* 用指针Swap */
	r2 := coordinate{5, 6}
	p := &r2
	p.ptrswap()
	fmt.Println(p, r2)

	var usr = struct {
		id       int
		name     string
		password string
	}{name: "zhangsan"}

	var usr1 = new(struct {
		id     int
		name   string
		passwd string
	})
	fmt.Println(usr, usr1)

	//嵌入式结构体用于map
	map1 := map[string]struct {
		name string
		age  int
	}{
		"teacher": {"zhenzhi", 39},
		"student": {"liming", 18},
	}
	fmt.Println(map1)

	map2 := new(map[string]struct {
		name string
		age  int
	})

	fmt.Println(map2)

	ds := dd{}
	ds1 := dd{d1{1}, d2{2}}

	fmt.Println(ds, ds1)

}
