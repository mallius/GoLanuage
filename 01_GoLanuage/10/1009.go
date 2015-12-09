package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("CPU number:", runtime.NumCPU())
	fmt.Println("Goroutines start:", runtime.NumGoroutine())
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println(n, runtime.NumGoroutine())
		}(i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Goroutines over:", runtime.NumGoroutine())
}
