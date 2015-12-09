package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.String())
	fmt.Println(t.Format("2006年01月02日"))
	fmt.Println(t.Weekday().String())
}
