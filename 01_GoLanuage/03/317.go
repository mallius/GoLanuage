package main

import (
	"fmt"
	"strconv"
)

func main() {
	var b bool = false
	var f float64 = 3.14
	var i int64 = -1024
	var ui uint64 = 1024

	fmt.Printf("bool[%t]->string[%q]\n", b, strconv.FormatBool(b))
	fmt.Printf("float[%f]->string[%q]\n", f, strconv.FormatFloat(f, 'f', 2, 32))
	fmt.Printf("int[%d]->string[%q]\n", i, strconv.FormatInt(i, 10))
	fmt.Printf("uint[%d]->string[%q]\n", ui, strconv.FormatUint(ui, 10))
}
