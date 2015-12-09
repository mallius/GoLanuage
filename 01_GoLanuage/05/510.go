package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num, col int
	var numStr string

	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			for k := 1; k < 5; k++ {
				if i != k && i != j && j != k {
					numStr = strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(k)
					num, _ = strconv.Atoi(numStr)
					fmt.Printf(" %d", num)
					col++
					if col == 5 {
						fmt.Printf("\n")
						col = 0
					}
				}
			}

		}
	}
}
