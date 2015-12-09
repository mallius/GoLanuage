package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			if i == 2 {
				runtime.Gosched()
			}
			fmt.Println("Goroutinel: ", i)
		}
	}()
	go func() {
		fmt.Println("Goroutine2")
	}()
	time.Sleep(5 * time.Second)
}
