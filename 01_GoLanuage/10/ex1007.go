package main

import (
	"fmt"
)

func main() {
	fmt.Println(fiboncii(1))
}

func fiboncii(n int) (ret chan int) {
	var n1, n2 chan int
	var ch chan int
	if n == 1 {
		ch <- 0
		close(ch)
	} else if n == 2 {
		ch <- 1
		close(ch)
	} else {
		n1 <- 0
		n2 <- 1
		for i := 3; i <= n; i++ {
			n2 <- ch
			n1 <- n2
		}
		close(ch)
	}

	return ch
}
