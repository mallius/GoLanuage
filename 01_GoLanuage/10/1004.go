package main

import (
	"fmt"
	"time"
)

func Worker(sem chan int, lock chan bool, id int) {
	sem <- 1 //相当于对同步信号量进行P原语操作
	fmt.Println(time.Now().Format("04:05"), id)
	time.Sleep(1 * time.Second)
	<-sem
	if id == 9 {
		lock <- true
	}
}
func main() {
	ch := make(chan int, 2)
	lock := make(chan bool)
	for i := 0; i < 10; i++ {
		go Worker(ch, lock, i)
	}
	<-lock
}
