package main

import (
	"fmt"
	"strconv"
)

func main() {
	var strb, strf, stri, strui string = "false", "3.14", "-1024", "1024"
	var b bool
	var f float64
	var i int64
	var ui uint64

	b, _ = strconv.ParseBool(strb)
	f, _ = strconv.ParseFloat(strf, 32)
	i, _ = strconv.ParseInt(stri, 10, 32)
	ui, _ = strconv.ParseUint(strui, 10, 32)

	fmt.Printf("string[%q]-->bool[%t]\n", strb, b)
	fmt.Printf("string[%q]-->bool[%f]\n", strf, f)
	fmt.Printf("string[%q]-->bool[%d]\n", stri, i)
	fmt.Printf("string[%q]-->bool[%d]\n", strui, ui)
}
