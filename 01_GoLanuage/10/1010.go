package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("Goroutinel defer...")
		for i := 0; i < 10; i++ {
			if i == 5 {
				runtime.Goexit()
			}
			fmt.Println("Goroutinel: ", i)
		}
	}()
	go func() {
		fmt.Println("Goroutinel2")
	}()
	time.Sleep(5 * time.Second)
}
