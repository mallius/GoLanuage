package main

import (
	"fmt"
	"math/rand"
)

func Test(ch chan int) {
	ch <- rand.Int()
	fmt.Println("Go...")
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Test(chs[i])
	}
	for _, ch := range chs {
		value := <-ch
		fmt.Println(value)
	}
}
