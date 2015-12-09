package main

import (
	"fmt"
)

func main() {
	var f float32 = 123.4567
	var cp = complex(1.2, 3.4)
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%6.2f\n", f)
	fmt.Printf("%v %g %g\n", cp, real(cp), imag(cp))

}
